<script setup>
</script>

<template>
    <div>
        <header class="bg-white shadow h-16 flex justify-between items-stretch">
            <router-link to="/" class="inline-block p-4 h-full hover:bg-teal-400 rounded hover:text-white font-bold">
                NavTrek
            </router-link>
            <div>
                <router-link :to="{name: 'tracking'}" class="inline-block p-4 h-full hover:bg-teal-400 rounded hover:text-white font-bold">
                    Tracking
                </router-link>
                <router-link :to="{name: 'login'}" class="inline-block p-4 h-full hover:bg-teal-400 rounded hover:text-white font-bold">
                    Login
                </router-link>
            </div>
        </header>
    </div>
</template>