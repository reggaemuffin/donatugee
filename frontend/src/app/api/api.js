import axios from 'axios';

export const HTTPS = axios.create({
  baseURL: `https://www.donatugee.de/api/v1/`,
});

export const createDonatorProfile = ({name, email}) => {
	return HTTPS.get('insert-donator', {
		params: {
			name,
			email,
		}
	}).then(response => response).catch(error => error.response);
};

export const createProfile = ({ name, email }) => {
	return HTTPS.get('insert-techfugee', {
		params: {
			name,
			email,
			skills: JSON.stringify([]),
		}
	}).then(response => response).catch(error => error.response);
};

export const loginDonator = ({ email }) => {
	return HTTPS.get('login-donator', {
		params: {
			email,
		}
	}).then(response => response).catch(error => error.response);
};

export const addSkills = ({ skills, id }) => {
	return HTTPS.get('add-skills', {
		params: {
			skills: JSON.stringify(skills),
			id,
		}
	}).then(response => response).catch(error => error.response);
};

export const getChallenges = () => {
  return HTTPS.get('challenges').then(response => response).catch(error => error.response);
};

export const getChallengeForTechfugee = (id) => {
	return HTTPS.get('application-by-techfugee', {
		params: {
			id
		}
	}).then(response => response).catch(error => error.response);
};

export const getChallenge = (id) => {
	return HTTPS.get('challenge', {
		params: {
			id
		}
	}).then(response => response).catch(error => error.response);
};

export const getTechfugee = (id) => {
	return HTTPS.get('techfugee', {
		params: {
			id
		}
	}).then(response => response).catch(error => error.response);
};

export const saveFurtherDetails = ({ id, city, introduction }) => {
	return HTTPS.get('update-techfugee', {
		params: {
			id,
			city,
			introduction
		}
	}).then(response => response).catch(error => error.response);
};


export const getDonator = (id) => {
	return HTTPS.get('donator', {
		params: {
			id,
		}
	}).then(response => response).catch(error => error.response);
};


export const getRandomText = (length) => {
	return axios.get('http://www.randomtext.me/api/gibberish/p-1/' + length).then(response => response).catch(error => error.response);
};

export const techfugeeAuthenticated = ({ id, passed }) => {
	return HTTPS.get('update-auth', {
		params: {
			id: id,
			passed: passed
		}
	}).then(response => response).catch(error => error.response)
};

export const setApplication = (challengeId, techfugeeId) => {
	return HTTPS.get('insert-application', {
		params: {
			'challenge_id': challengeId,
			'techfugee_id': techfugeeId
		}
	}).then(response => response).catch(error => error.response);
};

export const createCompanyProfile = ({ name, email, password, website, address }) => {
	return HTTPS.get('insert-donator', {
		params: {
			name,
			email,
			password,
			website,
			address,
		}
	}).then(response => response).catch(error => error.response);
};

export const getRefugee = ({id}) => {
	return HTTPS.get('techfugee', {
		params: {
			id,
		}
	}).then(response => response).catch(error => error.response);
};

export const retrieveChallengesForDonator = ({id}) => {
	return HTTPS.get('challenges-by-donator', {
		params: {
			id,
		}
	}).then(response => response).catch(error => error.response);
};

export const createChallengeForDonator = ({idDonator, name, description, duration, laptopType, amount, hardwareProvided}) => {
	return HTTPS.get('insert-challenge', {
		params: {
			id_donator: idDonator,
			name,
			description,
			duration,
			laptop_type: laptopType,
			amount,
			hardware_provided: hardwareProvided,
		}
	}).then(response => response).catch(error => error.response);
};

export const doAcceptApplicant = ({id}) => {
	return HTTPS.get('accept-application', {
		params: {
			id
		}
	}).then(response => response).catch(error => error.response);
};
