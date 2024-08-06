package handlers

import (
	"katsan/order/application/queries"
	"katsan/order/infrastructure/presentation/pages"
	"net/http"
)

func HandleGetAccountsPage(w http.ResponseWriter, r *http.Request) {
	response, err := queries.NewGetDomesticAccountsWithOrderStatusQueryHandler().Handle(queries.GetDomesticAccountsWithOrderStatusQuery{
		PageSize: func() string {
			if r.URL.Query().Get("pageSize") == "" {
				return "10"
			}
			return r.URL.Query().Get("pageSize")
		}(),
		PageOffset: func() string {
			if r.URL.Query().Get("pageOffset") == "" {
				return "0"
			}
			return r.URL.Query().Get("pageOffset")
		}(),
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	pages.AccountsPage(response).Render(r.Context(), w)
}

func HandleGetAccountDetailsAccountInformationPage(w http.ResponseWriter, r *http.Request) {
	pages.AccountDetailsAccountInformatinPage().Render(r.Context(), w)
}

func HandleGetAccountDetailsDebtAgingReportPage(w http.ResponseWriter, r *http.Request) {
	pages.AccountDetailsDebtAgingReportPage().Render(r.Context(), w)
}
