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
//		strncpy(stringarray[index], str->value, strlen(str->value));
//		index++;
//	}
//	rc_stringlist_free(stringlist);
//	stringarray[index] = NULL;
//	return stringarray;
//}
//void free_array(char** array) {
//	char **a;
//	for (a = array;  *a;  a++) {
//		if (*a == NULL) {
//			free(*a);
//			break;
//		}
//		free(*a);
//	}
//	free(array);
//}
//int rc_service_to_int (RC_SERVICE enu) {
//	return (int) enu;
//}
//char** getRunlevels() {
//	return rc_stringlist_to_char_array(rc_runlevel_list());
//}
//char** getServicesInRunlevel(const char* runlevel) {
//	return rc_stringlist_to_char_array(rc_services_in_runlevel(runlevel));
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
func GetRunlevels() []string {
	levels := C.getRunlevels()
	defer C.free_array(levels)
	ret := CStringArrayToGoStringArray(levels)
	return ret
}

// Get all services in specified runlevel
func GetServicesInRunlevel(runlevel string) []string {
	crunlevel := C.CString(runlevel)
	defer C.free(unsafe.Pointer(crunlevel))
	services := C.getServicesInRunlevel(crunlevel)
	defer C.free_array(services)
	ret := CStringArrayToGoStringArray(C.getServicesInRunlevel(crunlevel))
	return ret
}

// Get the encoded service state
func GetServiceState(service string) int {
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
	cservice := C.CString(service)
	defer C.free(unsafe.Pointer(cservice))
	state := int(C.getServiceState(cservice))
	return state
}
