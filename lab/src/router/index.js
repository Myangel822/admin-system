import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import TestView from '../views/TestView.vue'
import AchievementView from '../views/AchievementView.vue'
import PictureView from '../views/PictureView.vue'
import DemoView from '../views/DemoView.vue'
import TeamView from '../views/TeamView.vue'
import LoginView from '../views/LoginView.vue'
import RegisterView from '../views/RegisterView.vue'
import AdminView from '../views/AdminView.vue'
const routes = [
	// {
	//   path: '/',
	//   name: 'loginregister',
	//   // route level code-splitting
	//   // this generates a separate chunk (about.[hash].js) for this route
	//   // which is lazy-loaded when the route is visited.
	//   component: LoginRegisterView
	// },
	{
		path:'/',
		redirect:'/login'
	},
	{
		path:'/login',
		name:'login',
		component:LoginView
	},
	{
		path:'/register',
		name:'register',
		component:RegisterView
	},
  {
    path: '/home',
    name: 'home',
    component: HomeView
  },
  {
    path: '/test',
    name: 'test',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: TestView
  },
  {
    path: '/admin',
    name: 'admin',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: AdminView
  },
  {
    path: '/team',
    name: 'team',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: TeamView
  },
  {
    path: '/achievement',
    name: 'achievement',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: AchievementView
  },
  {
    path: '/demo',
    name: 'demo',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: DemoView
  },
  {
    path: '/picture',
    name: 'picture',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: PictureView
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
