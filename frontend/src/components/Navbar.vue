<script setup>
</script>

<template>
    <div>
        <header class="bg-white shadow h-16 flex justify-between items-stretch">
            <router-link to="/" class="inline-block p-4 h-full hover:bg-blue-200">
                One Step
            </router-link>
            <div>
                <router-link :to="{name: 'devices'}" class="inline-block p-4 h-full hover:bg-blue-200">
                    Devices
                </router-link>
                <router-link :to="{name: 'login'}" class="inline-block p-4 h-full hover:bg-blue-200">
                    Login
                </router-link>
            </div>
        </header>
    </div>
</template>