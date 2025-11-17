<script setup>
import { ref, onMounted } from 'vue'

const users = ref([])
const loading = ref(true)
const error = ref(null)
const showModal = ref(false)


// Fetch data from API
const fetchUsers = async () => {
  try {
    const response = await fetch('http://localhost:8080/users')
    if (!response.ok) {
      throw new Error('Failed to fetch data')
    }
    users.value = await response.json()
  } catch (err) {
    error.value = err.message
  } finally {
    loading.value = false
  }
}

const handleClick = () => {
  showModal.value = true
  console.log('Button clicked!')
}

const handleClose = () => {
  showModal.value = false
  console.log('Button clicked!')
}

// Fetch data when component mounts
onMounted(fetchUsers)
</script>

<template>
  <div class="flex flex-col h-screen items-center justify-start pt-10">
    <div class="pr-3 pb-3 w-140 flex justify-end">
      <button class="rounded bg-blue-200 p-2">Add</button>
    </div>
    <table class="table-auto shadow-lg bg-white rounded-md w-140">
      <thead class="bg-green-200">
        <tr>
          <th class="text-left p-2">ID</th>
          <th class="text-left p-2">ชื่อ สกุล</th>
          <th class="text-left p-2">ที่อยู่</th>
          <th class="text-left p-2">วันเกิด</th>
          <th class="text-left p-2">อายุ</th>
          <th class="text-left p-2">Action</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="user in users" :key="user.id">
          <td class="text-left p-2">{{ user.id }}</td>
          <td class="text-left p-2">{{ user.name }}</td>
          <td class="text-left p-2">{{ user.address }}</td>
          <td class="text-left p-2">{{ user.birthDate }}</td>
          <td class="text-left p-2">{{ user.age }}</td>
          <td class="p-2">
            <button @click="handleClick" class="rounded bg-blue-200 p-2">View</button>
          </td>
        </tr>
      </tbody>
    </table>
  </div>

  <!-- Modal -->
    <div
      v-if="showModal"
      class="fixed inset-0 flex items-center justify-center bg-opacity-30 backdrop-blur-sm"
    >
      <div class="bg-gray-100 rounded-lg w-96 p-6 relative">
        <h2 class="text-xl font-bold mb-4">Modal Title</h2>
        <p class="mb-4">This is a simple modal example using Vue 3 + TailwindCSS.</p>

        <!-- Close button -->
        <button
          @click="handleClose"
          class="absolute top-2 right-2 text-gray-500 hover:text-gray-800"
        >
          ✕
        </button>

        <button
          @click="confirmAction"
          class="bg-green-500 text-white px-4 py-2 rounded hover:bg-green-600"
        >
          Confirm
        </button>
      </div>
    </div>
</template>
