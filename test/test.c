#include "../src/Festi.h"

int main() {
	FS_Init();

	FS_Window* window = FS_CreateWindow(800, 600, "Test");
	while (!FS_WindowShouldClose())
    {	
    	FS_Refresh();
    }
	return 0;
}