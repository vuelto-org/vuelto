#include "../src/Vuelto.h"

int main() {
  vueltoInit();

  Vuelto_Window win = vueltoCreateWindow(800, 600, "test", true);

  Vuelto_Image img = vueltoLoadImage("test/coin.png", -0.2, -0.2, 0.7, 0.7);

  while (!vueltoWindowShouldClose(win)) {
    // vueltoDrawRect(-1, -1, 0.5, 0.5, 0.5, 0.1, 0.8);
    vueltoDrawImage(img);

    vueltoRefresh(win);
  }

  vueltoTerminate();
  return 0;
}
