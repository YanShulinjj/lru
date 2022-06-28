/* ----------------------------------
*  @author suyame 2022-06-28 9:58:00
*  Crazy for Golang !!!
*  IDE: GoLand
*-----------------------------------*/

package lru

import "errors"

var (
	KeyNotFoundError   = errors.New("Key Not Found !")
	KeyHasExistedError = errors.New("Key Has Existed!")
)
