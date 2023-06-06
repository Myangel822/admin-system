import { createStore } from 'vuex'

export default createStore({
  state: {
	  token:"",
	  userName:"",
  },
  getters: {
	  getToken(state){
		  return state.token ||localStorage.getTtem("token")||"";
		  }
	  },
	
  mutations: {
	  delToken(state){
	  		state.token = "";
	  		 localStorage.removeItem("token");
	  },
	  setToken(state,token){
		  state.token=token;
		  localStorage.setItem("token");
	  },
	  setUserInfo(state,userName){
		  state.userName=userName;
	  }
  },
  actions: {
  },
  modules: {
  }
})