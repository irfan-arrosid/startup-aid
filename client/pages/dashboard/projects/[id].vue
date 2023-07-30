<script setup>
import { ref } from 'vue'
const campaignId = ref(8)
const campaign = ref(null)

const response = await fetch(`http://localhost:8080/api/v1/campaigns/${campaignId.value}`)

const data = await response.json()
campaign.value = data.data

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
                    <h3 class="text-2xl text-gray-900 mb-4">Campaign Details</h3>
                </div>
                <div class="w-1/4 text-right">
                    <nuxt-link to="/dashboard/projects/create"
                        class="bg-green-button hover:bg-green-button text-white font-bold px-4 py-1 rounded inline-flex items-center">
                        Edit
                    </nuxt-link>
                </div>
            </div>
            <div class="block mb-2">
                <div class="w-full lg:max-w-full lg:flex mb-4">
                    <div
                        class="w-full border border-gray-400 bg-white rounded p-8 flex flex-col justify-between leading-normal">
                        <div>
                            <div class="text-gray-900 font-bold text-xl mb-2">
                                {{ campaign.name }}
                            </div>
                            <p class="text-sm font-bold flex items-center mb-1">
                                Description
                            </p>
                            <p class="text-gray-700 text-base">
                                {{ campaign.description }}
                            </p>
                            <p class="text-sm font-bold flex items-center mb-1 mt-4">
                                What Will Funders Get
                            </p>
                            <ul class="list-disc ml-5">
                                <li v-for="perk in campaign.perks">
                                    {{ perk }}
                                </li>
                            </ul>
                            <p class="text-sm font-bold flex items-center mb-1 mt-4">
                                Price
                            </p>
                            <p class="text-4xl text-gray-700">$ {{ new
                                Intl.NumberFormat().format(campaign.goal_amount) }}
                            </p>
                        </div>
                    </div>
                </div>
            </div>
            <div class="flex justify-between items-center">
                <div class="w-3/4 mr-6">
                    <h3 class="text-2xl text-gray-900 mb-4 mt-5">Gallery</h3>
                </div>
                <div class="w-1/4 text-right">
                    <a href="#"
                        class="bg-green-button hover:bg-green-button text-white font-bold px-4 py-1 rounded inline-flex items-center">
                        Upload
                    </a>
                </div>
            </div>
            <div class="flex -mx-2">
                <div class="relative w-1/4 bg-white m-2 p-2 border border-gray-400 rounded"
                    v-for="image in campaign.images">
                    <figure class="item-thumbnail">
                        <img src="{{ image }}" alt="" class="rounded w-full" />
                    </figure>
                </div>
            </div>
            <div class="flex justify-between items-center">
                <div class="w-3/4 mr-6">
                    <h3 class="text-2xl text-gray-900 mb-4 mt-5">
                        Transaction History
                    </h3>
                </div>
            </div>
            <div class="block mb-2">
                <div class="w-full lg:max-w-full lg:flex mb-4">
                    <div
                        class="w-full border border-gray-400 lg:border-gray-400 bg-white rounded p-8 flex flex-col justify-between leading-normal">
                        <div>
                            <div class="text-gray-900 font-bold text-xl mb-1">
                                Galih Pratama
                            </div>
                            <p class="text-sm text-gray-600 flex items-center mb-2">
                                Rp. 200.000 &middot; 12 September 2020
                            </p>
                        </div>
                    </div>
                </div>
            </div>
        </section>
        <div class="cta-clip -mt-20"></div>
        <section class="call-to-action bg-purple-progress pt-64 pb-10"></section>
        <Footer />
    </div>
</template>