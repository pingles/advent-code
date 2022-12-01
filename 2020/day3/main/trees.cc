#include <iostream>
#include "trees.h"

// modifies cursor according to the specified slope
std::pair<int, int> move(std::pair<int, int> cursor, int x, int y) {
  cursor.first += x;
  cursor.second += y;
  return cursor;
}

Item item_at(std::pair<int, int> cursor) {
  auto x = cursor.first % TERRAIN_COLS;
  auto y = cursor.second;
  return static_cast<Item>(terrain[y][x]);
}

// movement is a 2-item array, of x and y moves
int count_trees(const int movement[]) {
  auto position = std::pair<int, int>(0, 0);  // first = x, second = y

  position = move(position, movement[0], movement[1]);

  long number_of_trees = 0;

  // we loop until we reach the bottom
  while (position.second < TERRAIN_ROWS) {
    // the grid repeats to the right so we can mod 11 to loop around
    auto item = item_at(position);

    if (item == Item::TREE) {
      number_of_trees++;
    }

    // 3,1 slope
    position = move(position, movement[0], movement[1]);
  }

  return number_of_trees;
}

int main() {
  // strategies
  const int num_slopes = 5;
  const int slopes[num_slopes][2] = {
    {1,1},
    {3,1},
    {5,1},
    {7,1},
    {1,2},
  };

  long long product = 1;
  for (int i = 0; i < num_slopes; i++) {
    auto trees = count_trees(slopes[i]);
    product *= trees;
    std::cout << trees << std::endl;
  }
  std::cout << product << std::endl;

  return 0;
}
