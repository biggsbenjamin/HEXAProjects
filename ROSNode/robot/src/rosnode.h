#ifdef __cplusplus
extern "C" {
#endif

typedef struct {
    void * wrapper; // StringPublisherWrapper, because we can't use any c++ class directly in cgo.
}StringNode;

StringNode* NewStringNode(char *ip, char *subtopic, char *pubtopic);
void SubscribeString(StringNode* Node);
void DeleteStringNode(StringNode* Node);
void PublishString(StringNode* Node, char *data, int len);

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
