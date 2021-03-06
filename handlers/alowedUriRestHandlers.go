package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	m "github.com/Ulbora/GoAuth2/managers"
	oc "github.com/Ulbora/GoAuth2/oauthclient"
	"github.com/gorilla/mux"
)

/*
 Copyright (C) 2019 Ulbora Labs LLC. (www.ulboralabs.com)
 All rights reserved.

 Copyright (C) 2019 Ken Williamson
 All rights reserved.

 This program is free software: you can redistribute it and/or modify
 it under the terms of the GNU General Public License as published by
 the Free Software Foundation, either version 3 of the License, or
 (at your option) any later version.
 This program is distributed in the hope that it will be useful,
 but WITHOUT ANY WARRANTY; without even the implied warranty of
 MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 GNU General Public License for more details.
 You should have received a copy of the GNU General Public License
 along with this program.  If not, see <http://www.gnu.org/licenses/>.

*/

//AddAllowedURISuper AddAllowedURISuper
func (h *OauthRestHandler) AddAllowedURISuper(w http.ResponseWriter, r *http.Request) {
	//url of this endpoint
	var addsAuURL = "/ulbora/rs/clientAllowedUriSuper/add"

	var auscl oc.Claim
	auscl.Role = "superAdmin"
	auscl.URL = addsAuURL
	auscl.Scope = "write"
	h.Log.Debug("client: ", h.Client)
	auth := h.Client.Authorize(r, &auscl)
	if auth {
		h.SetContentType(w)
		aasURIContOk := h.CheckContent(r)
		h.Log.Debug("conOk: ", aasURIContOk)
		if !aasURIContOk {
			http.Error(w, "json required", http.StatusUnsupportedMediaType)
		} else {
			var cus m.ClientAllowedURI
			bsuc, berr := h.ProcessBody(r, &cus)
			h.Log.Debug("bsuc: ", bsuc)
			h.Log.Debug("cu: ", cus)
			h.Log.Debug("berr: ", berr)
			if !bsuc && berr != nil {
				http.Error(w, berr.Error(), http.StatusBadRequest)
			} else {
				ausSuc, ausID := h.Manager.AddClientAllowedURI(&cus)
				h.Log.Debug("auSuc: ", ausSuc)
				h.Log.Debug("auID: ", ausID)
				var rtn ResponseID
				if ausSuc && ausID != 0 {
					rtn.Success = ausSuc
					rtn.ID = ausID
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
				resJSON, _ := json.Marshal(rtn)
				fmt.Fprint(w, string(resJSON))
			}
		}
	} else {
		var fsrtn ResponseID
		w.WriteHeader(http.StatusUnauthorized)
		resJSON, _ := json.Marshal(fsrtn)
		fmt.Fprint(w, string(resJSON))
	}
}

//AddAllowedURI AddAllowedURI
func (h *OauthRestHandler) AddAllowedURI(w http.ResponseWriter, r *http.Request) {
	var cu m.ClientAllowedURI
	bsuc, berr := h.ProcessBody(r, &cu)
	h.Log.Debug("bsuc: ", bsuc)
	h.Log.Debug("cu: ", cu)
	h.Log.Debug("berr: ", berr)
	if bsuc && berr == nil {

		//url of this endpoint
		var addAuURL = "/ulbora/rs/clientAllowedUri/add"

		//make sure the user in not trying to add a prohibited url that has "ulbora" in the url
		//looks through a list of assets for the url and determines the role needed based on the asset part of the url
		acsuc, role := h.AssetControl.GetControlledAsset(addAuURL, "ulbora")

		var aucl oc.Claim
		if acsuc {
			aucl.Role = role
		} else {
			aucl.Role = "admin"
		}
		aucl.URL = addAuURL
		aucl.Scope = "write"
		h.Log.Debug("client: ", h.Client)

		//check that jwt token user role has permission to use the url of this endpoint
		auth := h.Client.Authorize(r, &aucl)

		if auth {
			// w.Header().Set("Content-Type", "application/json")
			h.SetContentType(w)
			aaURIContOk := h.CheckContent(r)
			h.Log.Debug("conOk: ", aaURIContOk)
			if !aaURIContOk {
				http.Error(w, "json required", http.StatusUnsupportedMediaType)
			} else {
				auSuc, auID := h.Manager.AddClientAllowedURI(&cu)
				h.Log.Debug("auSuc: ", auSuc)
				h.Log.Debug("auID: ", auID)
				var rtn ResponseID
				if auSuc && auID != 0 {
					rtn.Success = auSuc
					rtn.ID = auID
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
				resJSON, _ := json.Marshal(rtn)
				fmt.Fprint(w, string(resJSON))
			}
		} else {
			var frtn ResponseID
			w.WriteHeader(http.StatusUnauthorized)
			resJSON, _ := json.Marshal(frtn)
			fmt.Fprint(w, string(resJSON))
		}
	} else {
		http.Error(w, berr.Error(), http.StatusBadRequest)
	}
}

//UpdateAllowedURISuper UpdateAllowedURISuper
func (h *OauthRestHandler) UpdateAllowedURISuper(w http.ResponseWriter, r *http.Request) {
	//url of this endpoint
	h.Log.Debug("inside UpdateAllowedURISuper------------------------------------")

	var upsAuURL = "/ulbora/rs/clientAllowedUriSuper/update"

	var upuscl oc.Claim
	upuscl.Role = "superAdmin"
	upuscl.URL = upsAuURL
	upuscl.Scope = "write"
	h.Log.Debug("client: ", h.Client)
	auth := h.Client.Authorize(r, &upuscl)
	h.Log.Debug("auth: ", auth)
	if auth {
		// w.Header().Set("Content-Type", "application/json")
		h.SetContentType(w)
		uPasURIContOk := h.CheckContent(r)
		h.Log.Debug("conOk: ", uPasURIContOk)
		if !uPasURIContOk {
			http.Error(w, "json required", http.StatusUnsupportedMediaType)
		} else {
			var upcus m.ClientAllowedURI
			ubsuc, uberr := h.ProcessBody(r, &upcus)
			h.Log.Debug("ubsuc: ", ubsuc)
			h.Log.Debug("upcu: ", upcus)
			h.Log.Debug("uberr: ", uberr)
			if !ubsuc && uberr != nil {
				http.Error(w, uberr.Error(), http.StatusBadRequest)
			} else {
				uPusSuc := h.Manager.UpdateClientAllowedURI(&upcus)
				h.Log.Debug("auSuc: ", uPusSuc)
				var rtn Response
				if uPusSuc {
					rtn.Success = uPusSuc
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
				resJSON, _ := json.Marshal(rtn)
				fmt.Fprint(w, string(resJSON))
			}
		}
	} else {
		var fusrtn Response
		w.WriteHeader(http.StatusUnauthorized)
		resJSON, _ := json.Marshal(fusrtn)
		fmt.Fprint(w, string(resJSON))
	}
}

//UpdateAllowedURI UpdateAllowedURI
func (h *OauthRestHandler) UpdateAllowedURI(w http.ResponseWriter, r *http.Request) {
	var ucu m.ClientAllowedURI
	upbsuc, uberr := h.ProcessBody(r, &ucu)
	h.Log.Debug("upbsuc: ", upbsuc)
	h.Log.Debug("ucu: ", ucu)
	h.Log.Debug("uberr: ", uberr)
	if upbsuc && uberr == nil {

		//url of this endpoint
		var upAuURL = "/ulbora/rs/clientAllowedUri/update"

		//make sure the user in not trying to add a prohibited url that has "ulbora" in the url
		//looks through a list of assets for the url and determines the role needed based on the asset part of the url
		acsuc, role := h.AssetControl.GetControlledAsset(upAuURL, "ulbora")

		var aucl oc.Claim
		if acsuc {
			aucl.Role = role
		} else {
			aucl.Role = "admin"
		}
		aucl.URL = upAuURL
		aucl.Scope = "write"
		h.Log.Debug("client: ", h.Client)

		//check that jwt token user role has permission to use the url of this endpoint
		auth := h.Client.Authorize(r, &aucl)

		if auth {
			// w.Header().Set("Content-Type", "application/json")
			h.SetContentType(w)
			upaURIContOk := h.CheckContent(r)
			h.Log.Debug("conOk: ", upaURIContOk)
			if !upaURIContOk {
				http.Error(w, "json required", http.StatusUnsupportedMediaType)
			} else {
				uuSuc := h.Manager.UpdateClientAllowedURI(&ucu)
				h.Log.Debug("uuSuc: ", uuSuc)
				var rtn Response
				if uuSuc {
					rtn.Success = uuSuc
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
				resJSON, _ := json.Marshal(rtn)
				fmt.Fprint(w, string(resJSON))
			}
		} else {
			var frtn ResponseID
			w.WriteHeader(http.StatusUnauthorized)
			resJSON, _ := json.Marshal(frtn)
			fmt.Fprint(w, string(resJSON))
		}
	} else {
		http.Error(w, uberr.Error(), http.StatusBadRequest)
	}
}

//GetAllowedURI GetAllowedURI
func (h *OauthRestHandler) GetAllowedURI(w http.ResponseWriter, r *http.Request) {
	var getAuURL = "/ulbora/rs/clientAllowedUri/get"

	var guscl oc.Claim
	guscl.Role = "admin"
	guscl.URL = getAuURL
	guscl.Scope = "read"
	//fmt.Println("client: ", h.Client)
	auth := h.Client.Authorize(r, &guscl)
	if auth {
		//var id string
		h.SetContentType(w)
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) != 0 {
			var idStr = vars["id"]
			h.Log.Debug("vars: ", vars)
			id, idErr := strconv.ParseInt(idStr, 10, 64)
			if id != 0 && idErr == nil {
				h.Log.Debug("id: ", id)
				getAu := h.Manager.GetClientAllowedURI(id)
				h.Log.Debug("getAu: ", getAu)
				w.WriteHeader(http.StatusOK)
				resJSON, _ := json.Marshal(getAu)
				fmt.Fprint(w, string(resJSON))
			} else {
				w.WriteHeader(http.StatusBadRequest)
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

//GetAllowedURIList GetAllowedURIList
func (h *OauthRestHandler) GetAllowedURIList(w http.ResponseWriter, r *http.Request) {
	var getAulURL = "/ulbora/rs/clientAllowedUri/list"

	var gulcl oc.Claim
	gulcl.Role = "admin"
	gulcl.URL = getAulURL
	gulcl.Scope = "read"
	//fmt.Println("client: ", h.Client)
	auth := h.Client.Authorize(r, &gulcl)
	if auth {
		//var id string
		h.SetContentType(w)
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) != 0 {
			var clientIDStr = vars["clientId"]
			h.Log.Debug("vars: ", vars)
			clientID, idErr := strconv.ParseInt(clientIDStr, 10, 64)
			if clientID != 0 && idErr == nil {
				h.Log.Debug("clientID: ", clientID)
				getAul := h.Manager.GetClientAllowedURIList(clientID)
				h.Log.Debug("getAul: ", getAul)
				w.WriteHeader(http.StatusOK)
				resJSON, _ := json.Marshal(getAul)
				fmt.Fprint(w, string(resJSON))
			} else {
				w.WriteHeader(http.StatusBadRequest)
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

//DeleteAllowedURI DeleteAllowedURI
func (h *OauthRestHandler) DeleteAllowedURI(w http.ResponseWriter, r *http.Request) {
	var getAudURL = "/ulbora/rs/clientAllowedUri/delete"

	var gusdcl oc.Claim
	gusdcl.Role = "admin"
	gusdcl.URL = getAudURL
	gusdcl.Scope = "write"
	//fmt.Println("client: ", h.Client)
	auth := h.Client.Authorize(r, &gusdcl)
	if auth {
		//var id string
		h.SetContentType(w)
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) != 0 {
			var idStr = vars["id"]
			h.Log.Debug("vars delete: ", vars)
			id, idErr := strconv.ParseInt(idStr, 10, 64)
			h.Log.Debug("id delete: ", id)
			if id != 0 && idErr == nil {
				h.Log.Debug("id: ", id)
				getAud := h.Manager.DeleteClientAllowedURI(id)
				var rtn Response
				if getAud {
					rtn.Success = getAud
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
				resJSON, _ := json.Marshal(rtn)
				fmt.Fprint(w, string(resJSON))
			} else {
				w.WriteHeader(http.StatusBadRequest)
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}
