import { login } from '@/api/auth'
import { getToken, setToken, removeToken } from '@/utils/auth'

const user = {
  state: {
    token: getToken()
  },
  mutations: {
    SET_TOKEN (state, token) {
      state.token = token
    },
  },
  actions: {
    Login ({ commit }, userInfo) {
      // const username = userInfo.username.trim()
      return new Promise((resolve, reject) => {
        login(userInfo).then(response => {
          let data = response.data.data
          setToken(data)
          commit('SET_TOKEN', data)
          resolve()
        }).catch(error => {
          reject(error)
        })
      })
    }
  }
}
export default user
