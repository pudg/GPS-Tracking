<template>
	<main class="test w-full">
		<section class="forms bg-green-0 flex justify-center items-center">
            <div class="bg-white rounded-lg">
                <form class="login flex flex-col justify-center items-center h-full" @submit.prevent="handleLoginClick">
                    <h1 class="p-2 font-extrabold text-center">
						Login
					</h1>
                    <input
                        type="email" 
                        class="border m-1 text-center"
                        placeholder="Email address"
                        v-model="login_form.email" />
                    <input
                        class="border m-1 text-center"
                        type="password" 
                        placeholder="Password" 
                        v-model="login_form.password" />
                    <button
						class="font-extrabold my-2 p-4 text-lg text-indigo-100 transition-colors duration-150 bg-indigo-700 rounded-lg focus:shadow-outline hover:bg-indigo-800">
                        Login
                    </button>
					<p class="font-bold text-red-500">
						{{ loginError }}
					</p>
                </form>
            </div>
		</section>
	</main>
</template>

<script>
import { ref } from 'vue'
import { useStore, mapState, mapActions } from 'vuex'

export default {
	setup () {
		const login_form = ref({});
		const register_form = ref({});
		const store = useStore();

		const handleLoginClick = () => {
			store.dispatch('userAuthenticate', {
				email: login_form.value.email,
				password: login_form.value.password
			});
		};

		return {
			login_form,
			register_form,
			handleLoginClick,
		}
	},
	computed: {
		...mapState(['user', 'loginError']),
	},
	methods: {
		...mapActions(['userLogout']),
	}
}
</script>

<style>
.test {
	position: absolute;
	top: 40%;
}
.forms {
	display: flex;
	min-height: 25dvh;
}

form {
	flex: 1 1 0%;
	padding: 1rem;
}

form.register {
	color: #FFF;
	background-color: rgb(245, 66, 101);
	background-image: linear-gradient(
		to bottom right,
		rgb(245, 66, 101) 0%,
		rgb(189, 28, 60) 100%
	);
}

h2 {
	font-size: 2rem;
	text-transform: uppercase;
	margin-bottom: 2rem;
}

input {
	appearance: none;
	border: none;
	outline: none;
	background: none;

	display: block;
	/* width: 100%; */
	max-width: 400px;
	margin: 0 auto;
	font-size: 1.5rem;
	margin-bottom: 2rem;
	padding: 0.5rem 0rem;
}

input:not([type="submit"]) {
	opacity: 0.8;
	transition: 0.4s;
}

input:focus:not([type="submit"]) {
	opacity: 1;
}

input::placeholder {
	color: inherit;
}

form.register input:not([type="submit"]) {
	color: #FFF;
	border-bottom: 2px solid #FFF;
}

form.login input:not([type="submit"]) {
	color: #2c3e50;
	border-bottom: 2px solid #2c3e50;
}

form.login input[type="submit"] {
	background-color: rgb(245, 66, 101);
	color: #FFF;
	font-weight: 700;
	padding: 1rem 2rem;
	border-radius: 0.5rem;
	cursor: pointer;
	text-transform: uppercase;
}

form.register input[type="submit"] {
	background-color: #FFF;
	color: rgb(245, 66, 101);
	font-weight: 700;
	padding: 1rem 2rem;
	border-radius: 0.5rem;
	cursor: pointer;
	text-transform: uppercase;
}
</style>