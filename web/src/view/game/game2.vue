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

<script lang="ts" setup>
import { ref, onMounted } from 'vue';

const grid = ref([]);
const correctCount = ref(0);
const selectedCount = ref(0);
const cellClicked = ref<boolean[][]>([]);
const filledChars = ref<string[]>([]);
let correctChar = '';
const characters = ['戊', '戌', '乙']; // 文字种类


const generateGrid = () => {
  correctChar = characters[0]; // 设置第一个字符为正确的字符
  filledChars.value = characters; // 设置填充的字符
  for (let i = 0; i < 9; i++) {
    const row = [];
    const clickedRow = [];
    for (let j = 0; j < 9; j++) {
      const randomIndex = Math.floor(Math.random() * characters.length);
      row.push(characters[randomIndex]);
      clickedRow.push(false);
    }
    grid.value.push(row);
    cellClicked.value.push(clickedRow);
  }
};

const toggleCell = (rowIndex: number, colIndex: number) => {
  if (!cellClicked.value[rowIndex][colIndex]) {
    cellClicked.value[rowIndex][colIndex] = true;
    selectedCount.value++;
    if (grid.value[rowIndex][colIndex] === correctChar) {
      correctCount.value++;
    }
  } else {
    cellClicked.value[rowIndex][colIndex] = false;
    selectedCount.value--;
    if (grid.value[rowIndex][colIndex] === correctChar) {
      correctCount.value--;
    }
  }
};

onMounted(() => {
  generateGrid();

  // 在60秒后执行提交操作
  setTimeout(() => {
    submitSelection();
  }, 6000);

});


const submitSelection = () => {
  // 这里编写将选中的字符提交到api1接口的逻辑
  // 这里只是一个示例，你需要根据实际情况来编写提交的逻辑
  console.log('提交选中的字符到api1接口');
  console.log("=--cellClicked-",cellClicked)
  console.log("=--grid-",grid)
  console.log("=--characters-",characters)
  console.log("=--查找-",characters[0])

};

</script>

<style scoped>
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
