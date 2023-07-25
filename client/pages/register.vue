<script setup>
import { ref } from 'vue'

const name = ref('')
const occupation = ref('')
const email = ref('')
const password = ref('')

const register = async (name, occupation, email, password) => {
    const request = JSON.stringify({ name, occupation, email, password })

    try {
        const response = await fetch('http://localhost:8080/api/v1/users', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: request
        })
        const data = await response.json()

        if (response.ok) {
            useRouter().push({ path: '/upload' })
            console.log(data);
        } else {
            alert('sign up failed')
        }
    } catch (error) {
        console.error(error);
    }
}

</script>

<template>
    <div class="auth-page h-screen flex justify-center items-center">
        <div class="hidden md:block lg:w-1/3 bg-white h-full auth-background rounded-tr-lg rounded-br-lg"></div>
        <div class="w-auto md:w-2/4 lg:w-2/3 flex justify-center items-center">
            <div class="w-full lg:w-1/2 px-10 lg:px-0">
                <h2 class="font-normal mb-6 text-3xl text-white">
                    Sign Up Account
                </h2>
                <div class="mb-6">
                    <div class="mb-4">
                        <label class="font-normal text-lg text-white block mb-3">Full Name</label>
                        <input type="text" v-model="name"
                            class="auth-form focus:outline-none focus:bg-purple-hover focus:shadow-outline focus:border-purple-hover-stroke focus:text-gray-100"
                            placeholder="Write Your Name Here" />
                    </div>
                </div>
                <div class="mb-6">
                    <div class="mb-4">
                        <label class="font-normal text-lg text-white block mb-3">Occupation</label>
                        <input type="text" v-model="occupation"
                            class="auth-form focus:outline-none focus:bg-purple-hover focus:shadow-outline focus:border-purple-hover-stroke focus:text-gray-100"
                            placeholder="Write your occupation here" />
                    </div>
                </div>
                <div class="mb-6">
                    <div class="mb-4">
                        <label class="font-normal text-lg text-white block mb-3">Email Address</label>
                        <input type="email" v-model="email"
                            class="auth-form focus:outline-none focus:bg-purple-hover focus:shadow-outline focus:border-purple-hover-stroke focus:text-gray-100"
                            placeholder="Write your email address here" />
                    </div>
                </div>
                <div class="mb-6">
                    <div class="mb-4">
                        <label class="font-normal text-lg text-white block mb-3">Password</label>
                        <input type="password" v-model="password"
                            class="auth-form focus:outline-none focus:bg-purple-hover focus:shadow-outline focus:border-purple-hover-stroke focus:text-gray-100"
                            placeholder="Type your password here" />
                    </div>
                </div>
                <div class="mb-6">
                    <div class="mb-4">
                        <button @click="register(name, occupation, email, password)"
                            class="block w-full bg-orange-button hover:bg-green-button text-white font-semibold px-6 py-4 text-lg rounded-full">
                            Continue Sign Up
                        </button>
                    </div>
                </div>
                <div class="text-center">
                    <p class="text-white text-md">
                        Already have account?
                        <nuxt-link to="/login" class="no-underline text-orange-button">Sign In</nuxt-link>.
                    </p>
                </div>
            </div>
        </div>
    </div>
</template>

<style scoped>
.auth-background {
    background-image: url("/sign-up-background.jpg");
    background-position: center;
    background-size: cover;
}
</style>