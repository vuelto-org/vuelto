#include "renderer.hpp"

#include "../tools/definitions.hpp"
#include "app.hpp"

namespace Vuelto {
namespace Application {

Vuelto::Renderer2D CreateRenderer2D(Window win) {
  Vuelto::Renderer2D renderer;
  return renderer;
}

}  // namespace Application

const char* vertexShaderSource =
    "#version 330 core\n"
    "layout (location = 0) in vec3 aPos;\n"
    "void main()\n"
    "{\n"
    "   gl_Position = vec4(aPos.x, aPos.y, aPos.z, 1.0);\n"
    "}\0";

const char* fragmentShaderSource =
    "#version 330 core\n"
    "uniform vec3 color;\n"  // Added color uniform
    "out vec4 FragColor;\n"
    "void main()\n"
    "{\n"
    "   FragColor = vec4(color, 1.0f);\n"
    "}\n\0";

void Renderer2D::DrawRect(float x, float y, float width, float height, float r, float g, float b) {
  float vertices[] = {x, y, 0.0f, x + width, y, 0.0f, x + width, y + height, 0.0f, x, y + height, 0.0f};

  unsigned int VBO, VAO;
  glGenBuffers(1, &VBO);
  glGenVertexArrays(1, &VAO);

  glBindVertexArray(VAO);
  glBindBuffer(GL_ARRAY_BUFFER, VBO);
  glBufferData(GL_ARRAY_BUFFER, sizeof(vertices), vertices, GL_STATIC_DRAW);

  glVertexAttribPointer(0, 3, GL_FLOAT, GL_FALSE, 3 * sizeof(float), (void*)0);
  glEnableVertexAttribArray(0);

  unsigned int vertexShader, fragmentShader, shaderProgram;
  vertexShader = glCreateShader(GL_VERTEX_SHADER);
  glShaderSource(vertexShader, 1, &vertexShaderSource, NULL);
  glCompileShader(vertexShader);

  fragmentShader = glCreateShader(GL_FRAGMENT_SHADER);
  glShaderSource(fragmentShader, 1, &fragmentShaderSource, NULL);
  glCompileShader(fragmentShader);

  shaderProgram = glCreateProgram();
  glAttachShader(shaderProgram, vertexShader);
  glAttachShader(shaderProgram, fragmentShader);
  glLinkProgram(shaderProgram);
  glUseProgram(shaderProgram);

  glDeleteShader(vertexShader);
  glDeleteShader(fragmentShader);

  int colorLocation = glGetUniformLocation(shaderProgram, "color");
  glUniform3f(colorLocation, r, g, b);

  // Draw the rectangle
  glBindVertexArray(VAO);
  glDrawArrays(GL_TRIANGLE_FAN, 0, 4);

  // Cleanup
  glDeleteVertexArrays(1, &VAO);
  glDeleteBuffers(1, &VBO);
  glDeleteProgram(shaderProgram);
}

void Renderer2D::SetBackgroundColor(float color1, float color2, float color3) {
  glClearColor(color1, color2, color3, 1);
}

}  // namespace Vuelto
