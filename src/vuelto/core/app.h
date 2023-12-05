#include "../tools/definitions.h"

void vueltoInit();
void vueltoInitMultipleWindows();
Vuelto_Window vueltoCreateWindow(int width, int height, const char *title, bool resizable);
void vueltoDestroyWindow(Vuelto_Window win);
void vueltoTerminate();

bool vueltoWindowShouldClose(Vuelto_Window win);
void vueltoRefresh(Vuelto_Window win);
void vueltoMakeContextCurrent(Vuelto_Window win);
