package usecases

import (
	"Octoptimist/domain"
	"Octoptimist/tools"
	"strconv"
	"strings"
	"sync"
	"time"
)

type TribeManager struct {
}

func (tribeManager *TribeManager) Manage(
	peoples []domain.People,
	startDay time.Time,
	endDay time.Time,
	periodFiscal *domain.Period,
	bankHolydays *domain.BankHolidays,
	globalPurposeProjectsManager *domain.GlobalPurposeProjects,
	discountProjectsManager *domain.DiscountProjects,
	targetTaceManager *domain.TargetTaces,
	timeInputGetter domain.ITimeInputGetter,
	activityRateGetter domain.IActivityRateGetter,
	accessToken string,
	bypassNicknames string) []domain.PeopleInTribe {

	maxGoroutines := 8

	channelResult := make(chan domain.PeopleInTribe, len(peoples))

	var wg sync.WaitGroup

	for i := 0; i < len(peoples); i += maxGoroutines {
		limite := i + maxGoroutines
		if limite > len(peoples) {
			limite = len(peoples)
		}

		for j := i; j < limite; j++ {
			if strings.Contains(bypassNicknames, peoples[j].Nickname) {
				continue
			}

			if validEntryDate, err := time.Parse("2006-01-02", peoples[j].EntryDate); err == nil {
				if validEntryDate.After(endDay) {
					continue
				}
			}
			if validLeavingDate, err := time.Parse("2006-01-02", peoples[j].LeavingDate); err == nil {
				if validLeavingDate.Before(startDay) {
					continue
				}
			}

			wg.Add(1)
			go managePeople(
				peoples[j],
				startDay,
				endDay,
				*periodFiscal,
				bankHolydays,
				globalPurposeProjectsManager,
				discountProjectsManager,
				targetTaceManager,
				timeInputGetter,
				activityRateGetter,
				accessToken,
				time.Now(),
				channelResult,
				&wg)
		}

		wg.Wait()
	}

	close(channelResult)

	var peoplesInTribe []domain.PeopleInTribe
	for peopleInTribe := range channelResult {
		if !peopleInTribe.Actif {
			continue
		}
		peoplesInTribe = append(peoplesInTribe, peopleInTribe)
	}
	return peoplesInTribe
}

func managePeople(
	people domain.People,
	startDay time.Time,
	endDay time.Time,
	periodFiscal domain.Period,
	bankHolydays *domain.BankHolidays,
	globalPurposeProjectsManager *domain.GlobalPurposeProjects,
	discountProjectsManager *domain.DiscountProjects,
	targetTaceManager *domain.TargetTaces,
	timeInputGetter domain.ITimeInputGetter,
	activityRateGetter domain.IActivityRateGetter,
	accessToken string,
	pivot time.Time,
	channelResult chan domain.PeopleInTribe, wg *sync.WaitGroup) {
	defer wg.Done()

	peopleInTribe := domain.PeopleInTribe{}
	peopleInTribe.Actif = true
	peopleInTribe.Person = people

	var entryDate time.Time
	if convertedDay, err := time.Parse("2006-01-02", people.EntryDate); err == nil {
		entryDate = convertedDay
	}

	if periodFiscal.End.Before(entryDate) {
		peopleInTribe.Actif = false
		channelResult <- peopleInTribe
		return
	}

	if startDay.Before(entryDate) && entryDate.Before(endDay) {
		startDay = entryDate
	}

	if periodFiscal.Start.Before(entryDate) && entryDate.Before(periodFiscal.End) {
		periodFiscal.Start = entryDate
	}

	if validLeavingDay, err := time.Parse("2006-01-02", people.LeavingDate); err == nil {
		if (startDay.Before(validLeavingDay) || tools.DatesEquals(startDay, validLeavingDay)) && validLeavingDay.Before(endDay) {
			endDay = validLeavingDay
		}
		if periodFiscal.Start.Before(validLeavingDay) && validLeavingDay.Before(periodFiscal.End) {
			periodFiscal.End = validLeavingDay
		}
	}

	if targetTace, ok := targetTaceManager.GetTargetTaceForJobId(int(peopleInTribe.Person.JobId)); ok {
		peopleInTribe.TargetTace = targetTace
	}

	periodAnalysis := domain.NewPeriod(startDay, endDay, bankHolydays)

	totalWorkDays, err := periodAnalysis.TotalWorkDaysGetter()
	if err == nil {
		peopleInTribe.PeriodWorkDays = totalWorkDays
	}

	activityRate, err := activityRateGetter.Get(accessToken, strconv.FormatInt(people.ID, 10), tools.DateToString(periodAnalysis.Start), tools.DateToString(periodAnalysis.End))
	if err == nil {
		peopleInTribe.ActivityRates.OctopodFiscalYearActivityRate.Value = activityRate.Value
	}

	timeInput, err := timeInputGetter.Get(accessToken, strconv.FormatInt(people.ID, 10), tools.DateToString(periodAnalysis.Start), tools.DateToString(periodAnalysis.End), 50, globalPurposeProjectsManager)
	if err == nil {
		timeInput = timeInput.TimeInputEnricher(periodAnalysis, pivot)

		activityRatePeriodTace, billableTimeTotal, workDaysWithoutDayBreak := timeInput.ActivityRateCalculator(pivot, totalWorkDays)
		peopleInTribe.ActivityRates.RecalculatedPeriodActivityRate.Value = activityRatePeriodTace.Value
		peopleInTribe.TaceAble = workDaysWithoutDayBreak
		peopleInTribe.Taced = billableTimeTotal

		activityOptimistRateFiscalYear, billableOptimistTimeTotal, _ := timeInput.ActivityRateOptimistCalculator(pivot, totalWorkDays)
		peopleInTribe.ActivityRates.OptimistActivityRate.Value = activityOptimistRateFiscalYear.Value
		peopleInTribe.TacedOptimist = billableOptimistTimeTotal

		timeInputWithDiscount := timeInput.Clone()

		if timeInputWithDiscount != nil {
			timeInputWithDiscount = timeInputWithDiscount.TimeInputDiscountAdaptator(true, discountProjectsManager)

			activityRatePeriodWithDiscountTace, billableTimeTotalWithDiscount, _ := timeInputWithDiscount.ActivityRateCalculator(pivot, totalWorkDays)
			peopleInTribe.ActivityRates.RecalculatedPeriodWithDiscountActivityRate.Value = activityRatePeriodWithDiscountTace.Value
			peopleInTribe.TacedWithDiscount = billableTimeTotalWithDiscount

			activityRateOptimistWithDiscountPeriod, billableOptimistTimeTotalWithDiscount, _ := timeInputWithDiscount.ActivityRateOptimistCalculator(pivot, totalWorkDays)
			peopleInTribe.ActivityRates.OptimistWithDiscountActivityRate.Value = activityRateOptimistWithDiscountPeriod.Value
			peopleInTribe.TacedOptimistWithDiscount = billableOptimistTimeTotalWithDiscount
		}

	}

	channelResult <- peopleInTribe
}
