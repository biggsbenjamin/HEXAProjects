#ifdef __cplusplus
extern "C" {
#endif

typedef struct {
    void * wrapper; // StringPublisherWrapper, because we can't use any c++ class directly in cgo.
}StringSubscriber;

StringSubscriber* NewStringSubscriber(char *ip, char *topic);
void SubscribeString(StringSubscriber* sub);
void DeleteStringSubscriber(StringSubscriber* sub);

typedef struct {
    void * wrapper2; // StringPublisherWrapper, because we can't use any c++ class directly in cgo.
}StringPublisher;

StringPublisher* NewStringPublisher(char *ip, char *topic);
void PublishString(StringPublisher* pub, char *data, int len);
void DeleteStringPublisher(StringPublisher* pub);

#ifdef __cplusplus
}
#endif


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
