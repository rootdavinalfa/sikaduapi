/*
 * Copyright (c) 2019 - 2020. dvnlabs.xyz
 * Davin Alfarizky Putra Basudewa <dbasudewa@gmail.com>
 * API For sikadu.unbaja.ac.id
 */

package config

import "encoding/base64"

var TokenSecret = []byte("AjkasaL1_ow20jx_asiwALD2po_1sa-21")

func TokenSecretEncoded() []byte {
	sEnc := base64.StdEncoding.EncodeToString(TokenSecret)
	return []byte(sEnc)
}
