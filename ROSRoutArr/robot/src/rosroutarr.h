//sets up the wraper for using the c++ functions in go

#ifdef __cplusplus
extern "C" {
#endif

typedef struct {
    void * wrapper; // FloatSubscriberWrapper, because we can't use any c++ class directly in cgo.
}FloatSubscriber;

FloatSubscriber* NewFloatSubscriber(char *ip, char *topic);
void SpinSub(FloatSubscriber* sub);
void DeleteFloatSubscriber(FloatSubscriber* sub);

#ifdef __cplusplus
}
#endif


