<script setup>
import { ref, onMounted } from 'vue'
import axios from "axios";
import dayjs from 'dayjs';

const users = ref([])
const user = ref([])

const showCreateModal = ref(false)
const showViewModal = ref(false)

const userFormData = ref({
  name: "",
  surname:"",
  age:0,
  birthDate:"",
  address: ""
});

const fetchUsers = async () => {
  try {
    const response = await axios.get('http://localhost:8080/users')
    users.value = response.data
  } catch (err) {
    console.log("Error:", err);
  }
}

const submitUserFormData = async () => {
  showCreateModal.value = false
  try {
    userFormData.value.age = calculateAgeFromBirthDate(userFormData.value.birthDate)
    console.log("userFormData.value.age", userFormData.value.age)
    const response = await axios.post('http://localhost:8080/user', userFormData.value);
    console.log("Success:", response.data);
  } catch (err) {
    console.log("Error:", err);
  }
  fetchUsers()
}

const fetchUser = async (id) => {
  try {
    const response = await axios.get(`http://localhost:8080/user/${id}`,)    
    user.value = response.data
  } catch (err) {
    console.log("Error:", err);
  }
}

const handleOpenCreateModal = () => {
  showCreateModal.value = true
  console.log('Button clicked open create!')
}

const handleCloseCreateModal = () => {
  showCreateModal.value = false
  console.log('Button clicked close create!')
}

const handleOpenViewModal = (id) => {
  showViewModal.value = true
  fetchUser(id)
}

const handleCloseViewModal = () => {
  showViewModal.value = false
  console.log('Button clicked!')
}

function formatDate(dateString) {
  return dayjs(dateString).format("DD-MM-YYYY")
}

function calculateAgeFromBirthDate(birthDate) {
  if (birthDate === "") {
    return 0
  }
  var formatBirthDate = new Date(birthDate)
  var ageDifMs = Date.now() - formatBirthDate;
  var ageDate = new Date(ageDifMs);
  userFormData.value.age = Math.abs(ageDate.getUTCFullYear() - 1970);
  
  return userFormData.value.age
}
onMounted(fetchUsers())
</script>

<template>
  <div class="flex flex-col h-screen items-center justify-start pt-10">
    <div class="pr-3 pb-3 flex flex-col items-end w-200">
      <button @click="handleOpenCreateModal" class="rounded bg-blue-200 p-2">Add</button>
    </div>
    <table class="table-auto shadow-lg bg-white rounded-md w-200">
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
          <td class="text-left p-2">{{ user.name }} {{ user.surname }}</td>
          <td class="text-left p-2">{{ user.address }}</td>
          <td class="text-left p-2">{{ formatDate(user.birthDate) }}</td>
          <td class="text-left p-2">{{ user.age }}</td>
          <td class="p-2">
            <button @click="handleOpenViewModal(user.id)" class="rounded bg-blue-200 p-2">View</button>
          </td>
        </tr>
      </tbody>
    </table>
  </div>

  <!-- Modal -->
    <div
      v-if="showViewModal"
      class="fixed inset-0 flex items-center justify-center bg-opacity-30 backdrop-blur-sm"
    >
      <div class="flex flex-col justify-around rounded-lg p-10 w-200 h-100 border border-gray-200">
        <div class="flex flex-row w-150">
            <label for="name" class="flex items-center text-sm font-medium text-heading w-50">ชื่อ-สกุล</label>
            <div class="flex flex-row w-100 gap-4">
              <div id="name" class="w-50 border border-gray-300 rounded-md py-2 px-4" >{{ user.name }}</div>
              <div id="surname" class="w-50 border border-gray-300 rounded-md py-2 px-4">{{ user.surname }}</div>
            </div>
        </div>
        <div class="flex flex-row w-150">
          <label for="birth_date" class="flex items-center text-sm font-medium text-heading w-50">วันเกิด</label>
          <div id="birth_date" class="w-50 border border-gray-300 rounded-md py-2 px-4">{{ formatDate(user.birthDate) }}</div>
        </div>
        <div class="flex flex-row w-150">
            <label for="age" class="flex items-center text-sm font-medium text-heading  w-50">อายุ</label>
            <div id="age" class="w-50 border border-gray-300 rounded-md py-2 px-4">{{ user.age }}</div>
        </div>
        <div class="flex flex-row w-150">
            <label for="address" class="flex items-center text-sm font-medium text-heading  w-50">ที่อยู่</label>
            <div id="address" class="w-100 border border-gray-300 rounded-md py-2 px-4">{{ user.address }}</div>
        </div>

        <div class="flex flex-row w-full justify-end gap-4">
          <button
            @click="handleCloseViewModal"
            class="bg-red-500 text-white px-4 py-2 rounded hover:bg-red-600 w-20"
          >
            ยกเลิก
          </button>
        </div>
      </div>
    </div>

    <div
      v-if="showCreateModal"
      class="fixed inset-0 flex items-center justify-center bg-opacity-30 backdrop-blur-sm"
    >
      <div class="flex flex-col justify-around rounded-lg p-10 w-200 h-100 border border-gray-200">
        <div class="flex flex-row w-150">
            <label for="name" class="flex items-center text-sm font-medium text-heading w-50">ชื่อ-สกุล</label>
            <div class="flex flex-row w-100 gap-4">
              <input v-model="userFormData.name" type="text" id="name" class="border border-gray-300 rounded-md py-2 px-4 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent" required />
              <input v-model="userFormData.surname" type="text" id="surname" class="border border-gray-300 rounded-md py-2 px-4 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent" required />
            </div>
        </div>
        <div class="flex flex-row w-150">
           <label for="birth_date" class="flex items-center text-sm font-medium text-heading w-50">วันเกิด</label>
            <input v-model="userFormData.birthDate" type="date" id="birth_date" class="border border-gray-300 rounded-md py-2 px-4 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent" required />
        </div>
        <div class="flex flex-row w-150">
            <label for="age" class="flex items-center text-sm font-medium text-heading  w-50">อายุ</label>
            <div id="age" class="w-50 border border-gray-300 rounded-md py-2 px-4">{{ calculateAgeFromBirthDate(userFormData.birthDate) }}</div>
        </div>
        <div class="flex flex-row w-150">
            <label for="address" class="flex items-center text-sm font-medium text-heading w-50">ที่อยู่</label>
            <input v-model="userFormData.address" type="text" id="address" class="w-100 border border-gray-300 rounded-md py-2 px-4 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent" required />
        </div>

        <div class="flex flex-row w-full justify-end gap-4">
          <button
          @click="submitUserFormData"
          class="bg-green-500 text-white px-4 py-2 rounded hover:bg-green-600 w-20"
        >
          บันทึก
        </button>
        <button
          @click="handleCloseCreateModal"
          class="bg-red-500 text-white px-4 py-2 rounded hover:bg-red-600 w-20"
        >
          ยกเลิก
        </button>
        </div>
      </div>
    </div>
</template>
