#include "renderer.hpp"

#include <fstream>
#include <iostream>
#include <sstream>

#include "../tools/definitions.hpp"
#include "app.hpp"
#include "shader.hpp"

namespace Vuelto {
namespace Application {

Vuelto::Renderer2D CreateRenderer2D(Window win) {
  Vuelto::Renderer2D renderer;
  return renderer;
}

}  // namespace Application

const char* vertexShaderSource = Vuelto::Shader::ReadShaderFile("src/vuelto/shaders/vertexShader.glsl");

const char* fragmentShaderSource = Vuelto::Shader::ReadShaderFile("src/vuelto/shaders/fragmentShader.glsl");

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

void Renderer2D::DrawRectWithShader(float x, float y, float width, float height, const char* colorShader) {
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
  glShaderSource(fragmentShader, 1, &colorShader, NULL);
  glCompileShader(fragmentShader);

  shaderProgram = glCreateProgram();
  glAttachShader(shaderProgram, vertexShader);
  glAttachShader(shaderProgram, fragmentShader);
  glLinkProgram(shaderProgram);
  glUseProgram(shaderProgram);

  glDeleteShader(vertexShader);
  glDeleteShader(fragmentShader);

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
