//sets up the wraper for using the c++ functions in go

#ifdef __cplusplus
extern "C" {
#endif

typedef struct {
    void * wrapper; // FloatPublisherWrapper, because we can't use any c++ class directly in cgo.
}FloatSubscriber;

FloatSubscriber* NewFloatSubscriber(char *ip, char *topic);
void SubscribeFloat(FloatSubscriber* sub);
void DeleteFloatSubscriber(FloatSubscriber* sub);

#ifdef __cplusplus
}
#endif

//not sure what this bit is but it allows the variable to be manipulated across go and c++
#ifdef WIN32
#if defined(EXPORT_DLL)
#    define VAR __declspec(dllexport)
#elif defined(IMPORT_DLL)
#    define VAR __declspec(dllimport)
#endif
#else
#    define VAR extern
#endif
VAR char *subArray;
