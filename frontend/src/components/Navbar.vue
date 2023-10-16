<script setup>
</script>

<template>
    <div>
        <header class="bg-white shadow h-16 flex justify-between items-stretch">
            <router-link to="/" style="font-size: large;" class="inline-block p-4 h-full text-indigo-800 hover:bg-indigo-800 rounded hover:text-white font-extrabold">
                OneStep
            </router-link>
            <div>
                <router-link :to="{name: 'products'}" style="font-size: large;" class="inline-block p-4 h-full text-indigo-700 hover:bg-indigo-800 rounded hover:text-white font-extrabold">
                    TODO: Products
                </router-link>
                <router-link :to="{name: 'tracking'}" style="font-size: large;" class="inline-block p-4 h-full text-indigo-700 hover:bg-indigo-800 rounded hover:text-white font-extrabold">
                    Tracking
                </router-link>
                <router-link to="/" style="font-size: large;" class="inline-block p-4 h-full text-indigo-700 hover:bg-indigo-800 rounded hover:text-white font-extrabold">
                    TODO: Reviews
                </router-link>
            </div>
        </header>
    </div>
</template>