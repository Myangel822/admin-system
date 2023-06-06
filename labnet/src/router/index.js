import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

/* Layout */
import Layout from '@/layout'

/**
 * Note: sub-menu only appear when route children.length >= 1
 * Detail see: https://panjiachen.github.io/vue-element-admin-site/guide/essentials/router-and-nav.html
 *
 * hidden: true                   if set true, item will not show in the sidebar(default is false)
 * alwaysShow: true               if set true, will always show the root menu
 *                                if not set alwaysShow, when item has more than one children route,
 *                                it will becomes nested mode, otherwise not show the root menu
 * redirect: noRedirect           if set noRedirect will no redirect in the breadcrumb
 * name:'router-name'             the name is used by <keep-alive> (must set!!!)
 * meta : {
    roles: ['admin','editor']    control the page roles (you can set multiple roles)
    title: 'title'               the name show in sidebar and breadcrumb (recommend set)
    icon: 'svg-name'/'el-icon-x' the icon show in the sidebar
    breadcrumb: false            if set false, the item will hidden in breadcrumb(default is true)
    activeMenu: '/example/list'  if set path, the sidebar will highlight the path you set
  }
 */

/**
 * constantRoutes
 * a base page that does not have permission requirements
 * all roles can be accessed
 */
export const constantRoutes = [
  {
    path: '/login',
    component: () => import('@/views/login/index'),
    hidden: true
  },

  {
    path: '/404',
    component: () => import('@/views/404'),
    hidden: true
  },

  {
    path: '/',
    component: Layout,
    redirect: '/dashboard',
    children: [{
      path: 'dashboard',
      name: 'Dashboard',
      component: () => import('@/views/dashboard/index'),
      meta: { title: '首页', icon: 'dashboard' }
    }]
  },

  {
    path: '/list',
    component: Layout,
    name: 'List',
    meta: { title: '管理', icon: 'el-icon-s-help' },
    children: [
      {
        path: 'article',
        name: 'Article',
        component: () => import('@/views/table/index'),
        meta: { title: '论文', icon: 'table' }
      },
      {
        path: 'project',
        name: 'Project',
        component: () => import('@/views/table/project'),
        meta: { title: '项目', icon: 'table' }
      },
      {
        path: 'member',
        name: 'Member',
        component:() => import('@/views/table/member'),
        meta:{ title: '成员',icon: 'table'}
      },
      {
        path: 'standards',
        name: 'Standards',
        component:() => import('@/views/table/standards'),
        meta:{ title: '技术标准',icon: 'table'}
      },
      {
        path: 'activity',
        name: 'Activivy',
        component:() => import('@/views/table/activity'),
        meta:{ title: '活动',icon: 'table'}
      },
      {
        path: 'demo',
        name: 'Demo',
        component:() => import('@/views/table/demo'),
        meta:{ title: '演示',icon: 'table'}
      },
      {
        path: 'achievement',
        name: 'Achievement',
        component:() => import('@/views/table/achievement'),
        meta:{ title: '成就',icon: 'table'}
      },
      {
        path: 'news',
        name: 'News',
        component:() => import('@/views/table/news'),
        meta:{ title: '新闻',icon: 'table'}
      },
      {
        path: 'image',
        name: 'Image',
        component:() => import('@/views/table/image'),
        meta:{ title: '相册',icon: 'table'}
      }
    ]
  },

  {
    path: '/add',
    component:Layout,
    // redirect: '/add/article',
    name: 'Add',
    meta: { title: '表单', icon: 'el-icon-s-help' },
    children: [
      {
        path: 'index',
        name: 'AddArticle',
        component: () => import('@/views/form/index'),
        meta: { title: '论文', icon: 'form' }
      },
      {
        path: 'project',
        name: 'AddProject',
        component: () => import('@/views/form/project'),
        meta: { title: '项目', icon: 'form' }
      },
      {
        path: 'member',
        name: 'AddMember',
        component: () => import('@/views/form/member'),
        meta: { title: '成员', icon: 'form' }
      },
      {
        path: 'standards',
        name: 'AddStandards',
        component: () => import('@/views/form/standards'),
        meta: { title: '技术标准', icon: 'form' }
      },
      {
        path: 'activity',
        name: 'AddActivity',
        component: () => import('@/views/form/activity'),
        meta: { title: '活动', icon: 'form' }
      },
      {
        path: 'demo',
        name: 'AddDemo',
        component: () => import('@/views/form/demo'),
        meta: { title: '演示', icon: 'form' }
      },
      {
        path: 'achievement',
        name: 'AddAchievement',
        component: () => import('@/views/form/achievement'),
        meta: { title: '成就', icon: 'form' }
      },
      {
        path: 'news',
        name: 'AddNews',
        component: () => import('@/views/form/news'),
        meta: { title: '新闻', icon: 'form' }
      },
      {
        path: 'image',
        name: 'AddImage',
        component: () => import('@/views/form/image'),
        meta: { title: '相册', icon: 'form' }
      }
    ]
  },

  {
    path: '/edit',
    component: Layout,
    // redirect: '/add/article',
    name: 'Edit',

    children: [
      {
        path: '/index/:id',
        name: 'EditArticle',
        component: () => import('@/views/formEdit/index'),

      },
      {
        path: '/project/:id',
        name: 'EditProject',
        component: () => import('@/views/formEdit/project'),

      },
      {
        path: '/member/:id',
        name: 'EditMember',
        component: () => import('@/views/formEdit/member'),

      },
      {
        path: '/standards',
        name: 'EditStandards',
        component: () => import('@/views/formEdit/standards'),

      },
      {
        path: '/activity/:id',
        name: 'EditActivity',
        component: () => import('@/views/formEdit/activity'),

      },
      {
        path: '/demo/:id',
        name: 'EditDemo',
        component: () => import('@/views/formEdit/demo'),

      },
      {
        path: '/achievement/:id',
        name: 'EditAchievement',
        component: () => import('@/views/formEdit/achievement'),

      },
      {
        path: '/news/:id',
        name: 'EditNews',
        component: () => import('@/views/formEdit/news'),

      },
      {
        path: '/image/:id',
        name: 'EditImage',
        component: () => import('@/views/formEdit/image'),
      }
    ]
  },
  // 404 page must be placed at the end !!!
  { path: '*', redirect: '/404', hidden: true }
]

const createRouter = () => new Router({
  // mode: 'history', // require service support
  scrollBehavior: () => ({ y: 0 }),
  routes: constantRoutes
})

const router = createRouter()

// Detail see: https://github.com/vuejs/vue-router/issues/1234#issuecomment-357941465
export function resetRouter() {
  const newRouter = createRouter()
  router.matcher = newRouter.matcher // reset router
}

export default router
