import request from '@/utils/request'
import { getToken } from '@/utils/auth'

export function getArticle(token) {
  return request({
    url: '/api/get/articles',
    method: 'get',
    params: { token:getToken() }
  })
}

export function addArticle(data) {

  return request({
    url: '/api/v1/articles',
    method: 'post',
    params: {title:data.title,
      journal:data.journal,
      author:data.author,
      authors:data.authors,
      date:data.date,
      link:data.link,
      papercode:data.papercode,
      abstract:data.abstract,
      theyear:data.theyear,
      state:data.state,
      hide:data.hide,
      token:getToken() }
  })
}

export function getActivity(token){
  return request({
    url: '/api/get/activity',
    method: 'get',
    params: { token:getToken() }
  })
}
export function addActivity(data) {

  return request({
    url: '/api/v1/activity',
    method: 'post',
    params: {title:data.title,
      content:data.content,
      kind:data.kind,
      date:data.date,
      token:getToken() }
  })
}
export function getStandards(token){
  return request({
    url: '/api/get/standards',
    method: 'get',
    params: { token:getToken() }
  })
}
export function addStandards(data) {

  return request({
    url: '/api/v1/standards',
    method: 'post',
    params: {title:data.title,
      content:data.content,
      kind:data.kind,
      state:data.state,
      date:data.date,
      token:getToken() }
  })
}

export function deleteStandards(id){

  return request({
    url:'/api/v1/deleteStandard',
    method: 'delete',
    params:{
      id:id,
      token:getToken()}
  })
}

export function deleteArticle(id){

  return request({
    url:'/api/v1/deleteArticle',
    method: 'delete',
    params:{
      id:id,
      token:getToken()}
  })
}

export function deleteAchievement(id){

  return request({
    url:'/api/v1/deleteAchievement',
    method: 'delete',
    params:{
      id:id,
      token:getToken()}
  })
}
export function deleteProject(id){

  return request({
    url:'/api/v1/deleteProject',
    method: 'delete',
    params:{
      id:id,
      token:getToken()}
  })
}
export function deleteNews(id){

  return request({
    url:'/api/v1/deleteNews',
    method: 'delete',
    params:{
      id:id,
      token:getToken()}
  })
}
export function deleteMember(id){

  return request({
    url:'/api/v1/deleteMember',
    method: 'delete',
    params:{
      id:id,
      token:getToken()}
  })
}
export function deleteDemo(id){

  return request({
    url:'/api/v1/deleteDemo',
    method: 'delete',
    params:{
      id:id,
      token:getToken()}
  })
}
export function deleteActivity(id){

  return request({
    url:'/api/v1/deleteActivity',
    method: 'delete',
    params:{
      id:id,
      token:getToken()}
  })
}

export function getDemo(token){
  return request({
    url: '/api/get/demo',
    method: 'get',
    params: { token:getToken() }
  })
}
export function addDemo(data) {

  return request({
    url: '/api/v1/demo',
    method: 'post',
    params: {title:data.title,
      content:data.content,
      date:data.date,
      token:getToken() }
  })
}

export function getNews(token){
  return request({
    url: '/api/get/news',
    method: 'get',
    params: { token:getToken() }
  })
}
// export function addNews(title,content,imageUrl) {

//   return request({
//     url: '/api/v1/news',
//     method: 'post',
//     params: {
//       title:title,
//       content:content,
//       imageUrl:imageUrl,
//       token:getToken() }
//   })
// }

export function getAchievement(token){
  return request({
    url: '/api/get/achievements',
    method: 'get',
    params: { token:getToken() }
  })
}
export function addAchievement(data) {

  return request({
    url: '/api/v1/achievement',
    method: 'post',
    params: {title:data.title,
      kind:data.kind,
      content:data.content,
      date:data.date,
      token:getToken() }
  })
}



export function getMember(token){
  return request({
    url:'/api/get/members',
    method:'get',
    params:{token:getToken()}
  })
}
export function addMember(data){

    return request({
      url: '/api/v1/member',
      method: 'post',
      params:{
        username:data.username,
        password:data.password,
        name:data.name,
        phone:data.phone,
        graduation:data.graduation,
        graduationtime:data.graduationtime,
        mail:data.mail,
        identity:data.identity,
        introduction:data.introduction,
        research:data.research,
        achievement:data.achievement,
        state:data.state,
        token:getToken()
      }
    })
}


export function getProject(token) {
  return request({
    url: '/api/get/project',
    method: 'get',
    params: { token:getToken() }
  })
}

export function addProject(data) {

  return request({
    url: '/api/v1/project',
    method: 'post',
    params: {title:data.title,
      content:data.content,
      theyear:data.theyear,
      state:data.state,
      token:getToken() }
  })
}


export function getStandardsByName(name) {
  return request({
    url: '/api/get/mohuStandards',
    method: 'get',
    params:
    {
      name:name,
      token:getToken()
      }
  })
}

export function editStandards(id,data){
  return request({
    url:'/api/v1/editStandard',
    method:'put',
    params:
    {
      id:id,
      title:data.title,
      content:data.content,
      kind:data.kind,
      state:data.state,
      date:data.date,
      token:getToken()
    }
  })
}
export function editNews(id,data){
  return request({
    url:'/api/v1/editNews',
    method:'put',
    params:
    {
      id:id,
     title:data.title,
       content:data.content,

       token:getToken()
    }
  })
}

export function editProject(id,data){
  return request({
    url:'/api/v1/editProject',
    method:'put',
    params:
    {
      id:id,
     title:data.title,
       content:data.content,
       theyear:data.theyear,
       state:data.state,
       token:getToken()
    }
  })
}

export function editMember(id,data){
  return request({
    url:'/api/v1/editMember',
    method:'put',
    params:
    {
      id:id,
     name:data.name,
     phone:data.phone,
     mail:data.mail,
     graduation:data.graduation,
     graduationtime:data.graduationtime,
     identity:data.identity,
     introduction:data.introduction,
     research:data.research,
     achievement:data.achievement,
     state:data.state,
     token:getToken()
    }
  })
}

export function editArticle(data){
  return request({
    url:'/api/v1/editArticle',
    method:'put',
    params:
    {
      id:data.id,
    title:data.title,
      journal:data.journal,
      author:data.author,
      authors:data.authors,
      date:data.date,
      link:data.link,
      papercode:data.papercode,
      abstract:data.abstract,
      theyear:data.theyear,
      state:data.state,
      isHide:data.isHide,
      token:getToken()
    }
  })
}

export function editDemo(id,data){
  return request({
    url:'/api/v1/editDemo',
    method:'put',
    params:
    {
      id:id,
    title:data.title,
      content:data.content,
      date:data.date,
      token:getToken()
    }
  })
}

export function editActivity(id,data){
  return request({
    url:'/api/v1/editActivity',
    method:'put',
    params:
    {
      id:id,
   title:data.title,
     content:data.content,
     kind:data.kind,
     date:data.date,
     token:getToken()
    }
  })
}


export function editAchievement(id,data){
  return request({
    url:'/api/v1/editAchievement',
    method:'put',
    params:
    {
      id:id,
   title:data.title,
     kind:data.kind,
     content:data.content,
     date:data.date,
     token:getToken()
    }
  })
}

export function resetPassword(id){
  return request({
    url:'/api/v1/resetPassword',
    method:'put',
    params:
    {
      id:id,
      token:getToken()
    }
  })
}

export function getProjectByName(name) {
  return request({
    url: '/api/get/mohuProject',
    method: 'get',
    params:
    {
      name:name,
      token:getToken()
      }
  })
}

export function getNewsByName(name) {
  return request({
    url: '/api/get/mohuNews',
    method: 'get',
    params:
    {
      name:name,
      token:getToken()
      }
  })
}

export function getMemberByName(name) {
  return request({
    url: '/api/get/mohuMember',
    method: 'get',
    params:
    {
      name:name,
      token:getToken()
      }
  })
}

export function getArticleByName(name) {
  return request({
    url: '/api/get/mohuArticle',
    method: 'get',
    params:
    {
      name:name,
      token:getToken()
      }
  })
}

export function getDemoByName(name) {
  return request({
    url: '/api/get/mohuDemo',
    method: 'get',
    params:
    {
      name:name,
      token:getToken()
      }
  })
}

export function getActivityByName(name) {
  return request({
    url: '/api/get/mohuActivity',
    method: 'get',
    params:
    {
      name:name,
      token:getToken()
      }
  })
}

export function getAchievementByName(name) {
  return request({
    url: '/api/get/mohuAchievement',
    method: 'get',
    params:
    {
      name:name,
      token:getToken()
      }
  })
}

export function getActivityByKind(kind) {
  return request({
    url: '/api/get/getActivityByKind',
    method: 'get',
    params:
    {
      kind:kind,
      token:getToken()
      }
  })
}

export function countArticle() {
  return request({
    url: '/api/get/countArticle',
    method: 'get',
    params:
    {
      token:getToken()
      }
  })
}

export function countMember() {
  return request({
    url: '/api/get/countMember',
    method: 'get',
    params:
    {
      token:getToken()
      }
  })
}

export function countProject() {
  return request({
    url: '/api/get/countProject',
    method: 'get',
    params:
    {
      token:getToken()
      }
  })
}


export function editStandardState(data) {
  return request({
    url: '/api/v1/editStandardState',
    method: 'put',
    params:
    {
      state:data.state,
      id:data.id,
      token:getToken()
      }
  })
}

export function editProjectState(data) {
  return request({
    url: '/api/v1/editProjectState',
    method: 'put',
    params:
    {
      state:data.state,
      id:data.id,
      token:getToken()
      }
  })
}

export function editArticleState(data) {
  return request({
    url: '/api/v1/editArticleState',
    method: 'put',
    params:
    {
      state:data.state,
      id:data.id,
      token:getToken()
      }
  })
}

export function editArticleHide(data) {
  return request({
    url: '/api/v1/editArticleHide',
    method: 'put',
    params:
    {
      hide:data.hide,
      id:data.id,
      token:getToken()
      }
  })
}

export function getImages(){
  return request({
    url:'/api/get/images',
    method:'get',
    params:
    {
      token:getToken()
    }
  })
}

export function getImagesByDate(date){
  return request({
    url:'/api/get/mohuImages',
    method:'get',
    params:
    {
      date:date,
      token:getToken()
    }
  })
}

export function deleteImage(id){
  return request({
    url:'/api/v1/deleteImage',
    method:'delete',
    params:
    {
      id:id,
      token:getToken()
    }
  })
}




