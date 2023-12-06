#include "../tools/definitions.h"
#include "structs.h"

void vueltoDrawRect(float x, float y, float width, float height, float r, float g, float b);
void vueltoSetBackgroundColor(float color1, float color2, float color3);
void vueltoDrawLine(int x1, int x2, int y1, int y2, float color1, float color2, float color3);

Vuelto_Image vueltoLoadImage(const char* imagePath, float x, float y, float width, float height);
void vueltoDrawImage(Vuelto_Image img);

void vueltoCleanUp();
