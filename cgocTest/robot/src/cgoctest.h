//establishing c++ function for use in go
#ifdef __cplusplus
extern "C" {
#endif

int main2();

#ifdef __cplusplus
}
#endif

//got this code from the golang website
#ifdef WIN32
#if defined(EXPORT_DLL)
#    define VAR __declspec(dllexport)
#elif defined(IMPORT_DLL)
#    define VAR __declspec(dllimport)
#endif
#else
#    define VAR extern
#endif
VAR char *hello;
