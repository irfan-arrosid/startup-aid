<script setup>
import { ref } from 'vue'

const name = ref('')
const short_description = ref('')
const description = ref('')
const goal_amount = ref()
const perks = ref('')

const userId = ref(10)
const token = ref('eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxMH0.8uhX9oAbQ4_Pvd6zr97l05KyIVksoeujG8uH3LR5jVQ')

const createCampaign = async (name, short_description, description, goal_amount, perks) => {
    const request = JSON.stringify({ name, short_description, description, goal_amount, perks })

    try {
        const response = await fetch(`http://localhost:8080/api/v1/campaigns?user_id=${userId.value}`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': `Bearer ${token.value}`
            },
            body: request
        })

        const data = await response.json()

        if (response.ok) {
            useRouter().push({ path: '/dashboard' })
            console.log(data);
        } else {
            alert('Create campaign is failed')
        }
    } catch (error) {
        console.error(error)
    }
}

</script>

<template>
    <div class="project-page">
        <section class="dashboard-header pt-5">
            <div class="container mx-auto relative">
                <Navbar />
            </div>
        </section>
        <section class="container mx-auto pt-8">
            <div class="flex justify-between items-center">
                <div class="w-full mr-6">
                    <h2 class="text-4xl text-gray-900 mb-2 font-medium">Dashboard</h2>
                </div>
            </div>
            <div class="flex justify-between items-center">
                <div class="w-3/4 mr-6">
                    <h3 class="text-2xl text-gray-900 mb-4">Create New Projects</h3>
                </div>
                <div class="w-1/4 text-right">
                    <button @click="createCampaign(name, short_description, description, goal_amount, perks)"
                        class="bg-green-button hover:bg-green-button text-white font-bold px-4 py-1 rounded inline-flex items-center">
                        Save
                    </button>
                </div>
            </div>
            <div class="block mb-2">
                <div class="w-full lg:max-w-full lg:flex mb-4">
                    <div
                        class="w-full border border-gray-400 bg-white rounded p-8 flex flex-col justify-between leading-normal">
                        <form class="w-full">
                            <div class="flex flex-wrap -mx-3 mb-6">
                                <div class="w-full md:w-1/2 px-3 mb-6 md:mb-0">
                                    <label class="block uppercase tracking-wide text-gray-700 text-xs font-bold mb-2">
                                        Campaign Name
                                    </label>
                                    <input
                                        class="appearance-none block w-full bg-gray-200 text-gray-700 border border-gray-200 rounded py-3 px-4 leading-tight focus:outline-none focus:bg-white focus:border-gray-500"
                                        type="text" placeholder="Contoh: Demi Gunpla Demi Istri" v-model="name" />
                                </div>
                                <div class="w-full md:w-1/2 px-3">
                                    <label class="block uppercase tracking-wide text-gray-700 text-xs font-bold mb-2">
                                        Price
                                    </label>
                                    <input
                                        class="appearance-none block w-full bg-gray-200 text-gray-700 border border-gray-200 rounded py-3 px-4 leading-tight focus:outline-none focus:bg-white focus:border-gray-500"
                                        type="number" placeholder="Contoh: 200000" v-model="goal_amount" />
                                </div>
                                <div class="w-full px-3">
                                    <label class="block uppercase tracking-wide text-gray-700 text-xs font-bold mb-2 mt-3">
                                        Short Description
                                    </label>
                                    <input
                                        class="appearance-none block w-full bg-gray-200 text-gray-700 border border-gray-200 rounded py-3 px-4 mb-3 leading-tight focus:outline-none focus:bg-white focus:border-gray-500"
                                        type="text" placeholder="Deskripsi singkat mengenai projectmu"
                                        v-model="short_description" />
                                </div>
                                <div class="w-full px-3">
                                    <label class="block uppercase tracking-wide text-gray-700 text-xs font-bold mb-2">
                                        What will backers get
                                    </label>
                                    <input
                                        class="appearance-none block w-full bg-gray-200 text-gray-700 border border-gray-200 rounded py-3 px-4 mb-3 leading-tight focus:outline-none focus:bg-white focus:border-gray-500"
                                        type="text" placeholder="Contoh: Ayam, Nasi Goreng, Piring" v-model="perks" />
                                </div>
                                <div class="w-full px-3">
                                    <label class="block uppercase tracking-wide text-gray-700 text-xs font-bold mb-2">
                                        Description
                                    </label>
                                    <textarea
                                        class="appearance-none block w-full bg-gray-200 text-gray-700 border border-gray-200 rounded py-3 px-4 mb-3 leading-tight focus:outline-none focus:bg-white focus:border-gray-500"
                                        type="text" placeholder="Isi deskripsi panjang untuk projectmu"
                                        v-model="description"></textarea>
                                </div>
                            </div>
                        </form>
                    </div>
                </div>
            </div>
        </section>
        <div class="cta-clip -mt-20"></div>
        <section class="call-to-action bg-purple-progress pt-64 pb-10"></section>
        <Footer />
    </div>
</template>