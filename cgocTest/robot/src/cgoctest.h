#ifdef __cplusplus
extern "C" {
#endif

int main2();

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
VAR char *hello;
