#include "../src/Festi.h"

int main() {
	FS_Init();

	FS_Window* window = FS_CreateWindow(800, 600, "Test");

	while (!FS_WindowShouldClose())
    {	
    	static float color_bc[3] = {0.5, 0.5, 0.5};
    	static float color[3] = {0.1, 0.1, 0.1};



    	FS_SetBackgroundColor(color_bc);

    	FS_DrawRect(0, 0, 0.5, 0.5, color);

    	FS_Refresh();
    }
	return 0;
}
