<template>
    <div>
        <v-flex xs12>
            <h1>
                Log In
            </h1>
            <v-alert v-if="errorMessage !== ''" color="error" icon="warning" value="true">
                {{ this.errorMessage }}
            </v-alert>
            <v-form v-model="valid">
                <v-text-field
                        label="E-mail"
                        v-model="email"
                        :rules="emailRules"
                        required
                ></v-text-field>
                <v-text-field
                        label="Password"
                        v-model="password"
                        :rules="passwordRules"
                        type="password"
                        required
                ></v-text-field>
                <v-btn class="sign-up-button"
                       small
                       color="primary"
                       dark
                       large
                       @click="signUp"
                >
                    Log In
                </v-btn>
            </v-form>
        </v-flex>
    </div>
</template>
<script>
	import { createProfile } from '../../api/api';

	export default {
		name: 'Login',
		data () {
			return {
				valid: false,
				name: '',
				nameRules: [
					(v) => !!v || 'Name is required',
				],
				email: '',
				emailRules: [
					(v) => !!v || 'E-mail is required',
					(v) => /^\w+([\.-]?\w+)*@\w+([\.-]?\w+)*(\.\w{2,3})+$/.test(v) || 'E-mail must be valid'
				],
				errorMessage: '',
				password: '',
				passwordRules: [
					(v) => !!v || 'Password is required',
				],
			}
		},
		computed: {
			disabled() {
				return !this.valid;
			}
		},
		methods: {
			signUp() {
				createProfile({
					name: this.name,
					email: this.email,
					password: this.password
				}).then(response => {
					if (response.status !== 200) {
						this.errorMessage = 'User cannot be created';
					}

					let authenticated;
					if (response.data.Authenticated === '') {
                        authenticated = false
                    } else {
						authenticated = JSON.parse(response.data.Authenticated);
                    }

					window.localStorage.setItem('userId', response.data.ID);
					window.localStorage.setItem('wrongAnswers', 0);

					if (Object.keys(JSON.parse(window.localStorage.getItem('skills'))).length === 0) {
						return this.$router.push({
							path: '/interests/add/' + response.data.ID,
						});
                    }

					if (!authenticated) {
						return this.$router.push({
							path: '/tech-questions/1',
						});
                    }

					if (!response.data.Introduction === '' || response.data.City === '') {
						return this.$router.push({
							path: '/refugee/further-details',
						});
					}

					const idChallenge = window.localStorage.getItem('idChallenge');
					if (idChallenge === null) {
						this.$router.push({
                            path: '/challenges',
                        });
                    }

					return this.$router.push({
						path: '/challenge/' + idChallenge,
					});
				})
			}
		}
	};
</script>

<style lang="scss" type="text/scss">
    .sign-up-button {
        width: 100%;
    }
</style>