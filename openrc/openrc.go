package openrc

// #cgo LDFLAGS: -lrc
// #include <stdlib.h>
// #include <string.h>
// #include <rc.h>
//
//char** rc_stringlist_to_char_array(RC_STRINGLIST *stringlist) {
//	char **stringarray;
// 	RC_STRING *str;
//	int index = 0;
//	for ((str) = ((stringlist)->tqh_first); 	\
//		(str) != (NULL);			\
//		(str) = ((str)->entries.tqe_next)) {
//		index++;
//	}
//	stringarray =  (char**) calloc(sizeof(char*), index + 1);
//	index = 0;
//	for ((str) = ((stringlist)->tqh_first); 	\
//		(str) != (NULL);			\
//		(str) = ((str)->entries.tqe_next)) {
//		stringarray[index] = (char*) calloc(sizeof(char), strlen(str->value) + 1);
//		strcpy(stringarray[index], str->value);
//		index++;
//	}
//	return stringarray;
//}
//int rc_service_to_int (RC_SERVICE enu) {
//	return (int) enu;
//}
//char** getRunlevels() {
//		RC_STRINGLIST *levelsl;
//		levelsl = rc_runlevel_list();
//		return rc_stringlist_to_char_array(levelsl);
//}
//char** getServicesInRunlevel(const char* runlevel) {
//		RC_STRINGLIST *servicesl;
//		servicesl = rc_services_in_runlevel(runlevel);
//		return rc_stringlist_to_char_array(servicesl);
//}
//int getServiceState(const char* service) {
//	return rc_service_to_int(rc_service_state(service));
//}
import "C"
import (
	"unsafe"
)

// Converts char** to []string
func CStringArrayToGoStringArray(p **_Ctype_char) []string {
	var strings []string
	q := uintptr(unsafe.Pointer(p))
	for {
		p = (**C.char)(unsafe.Pointer(q))
		if *p == nil {
			break
		}
		strings = append(strings, C.GoString(*p))
		q += unsafe.Sizeof(q)
	}
	return strings
}

// Get all runlevels
func getRunlevels() []string {
	levels := C.getRunlevels()
	return CStringArrayToGoStringArray(levels)
}

// Get all services in specified runlevel
func getServicesInRunlevel(runlevel string) []string {
	services := C.getServicesInRunlevel(C.CString(runlevel))
	return CStringArrayToGoStringArray(services)
}

// Get the encoded service state
func getServiceState(service string) int {
	/*! @brief States a service can be in
	typedef enum
	{
		These are actual states
		The service has to be in one only at all times
		RC_SERVICE_STOPPED     = 0x0001,
		RC_SERVICE_STARTED     = 0x0002,
		RC_SERVICE_STOPPING    = 0x0004,
		RC_SERVICE_STARTING    = 0x0008,
		RC_SERVICE_INACTIVE    = 0x0010,

	Service may or may not have been hotplugged
		RC_SERVICE_HOTPLUGGED = 0x0100,

	Optional states service could also be in
		RC_SERVICE_FAILED      = 0x0200,
		RC_SERVICE_SCHEDULED   = 0x0400,
		RC_SERVICE_WASINACTIVE = 0x0800
	} RC_SERVICE;
	*/
	return int(C.getServiceState(C.CString(service)))
}
