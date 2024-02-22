<template>
  <div id="app">
    <div class="grid">
      <div v-for="(row, rowIndex) in grid" :key="rowIndex" class="row">
        <div v-for="(cell, colIndex) in row" :key="colIndex" class="cell" @click="toggleCell(rowIndex, colIndex)" :class="{ 'green-bg': cellClicked[rowIndex][colIndex] }">
          {{ cell }}
        </div>
      </div>
    </div>
    <div class="score">
      <p>正确个数: {{ correctCount }}</p>
      <p>选中个数: {{ selectedCount }}</p>
    </div>
    <div class="example">
      <p>请找出字符: {{ filledChars[0] }}</p>
     <!-- <p>填充的字符: {{ filledChars.join(', ') }}</p>-->
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      grid: [],
      correctCount: 0,
      selectedCount: 0,
      cellClicked: [], // 记录每个格子是否被点击过
      filledChars: [], // 填充的字符
      correctChar: '' // 正确的字符
    };
  },
  mounted() {
    this.generateGrid();
  },
  methods: {
    generateGrid() {
      const characters = ['戊', '戌', '乙']; // 文字种类
      this.correctChar = characters[0]; // 设置第一个字符为正确的字符
      this.filledChars = characters; // 设置填充的字符
      for (let i = 0; i < 9; i++) {
        const row = [];
        const clickedRow = [];
        for (let j = 0; j < 9; j++) {
          const randomIndex = Math.floor(Math.random() * characters.length);
          row.push(characters[randomIndex]);
          clickedRow.push(false);
        }
        this.grid.push(row);
        this.cellClicked.push(clickedRow);
      }
    },
    toggleCell(rowIndex, colIndex) {
      if (!this.cellClicked[rowIndex][colIndex]) {
        this.cellClicked[rowIndex][colIndex] = true;
        this.selectedCount++;
        if (this.grid[rowIndex][colIndex] === this.correctChar) {
          this.correctCount++;
        }
      } else {
        this.cellClicked[rowIndex][colIndex] = false;
        this.selectedCount--;
        if (this.grid[rowIndex][colIndex] === this.correctChar) {
          this.correctCount--;
        }
      }
    }
  }
};
</script>

<style>
.grid {
  display: flex;
  flex-direction: column;
  border: 1px solid #000;
  margin-bottom: 20px;
}

.row {
  display: flex;
}

.cell {
  width: 50px;
  height: 50px;
  border: 1px solid #000;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
}

.green-bg {
  background-color: green;
}
</style>

