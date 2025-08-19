import { createRouter, createWebHistory } from 'vue-router'
import WizardView from '../views/WizardView.vue'
import LoginView from '../views/LoginView.vue'
import OnboardingView from '../views/OnboardingView.vue'
import ApiTestView from '../views/ApiTestView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    { path: '/', redirect: '/onboarding' },
    { path: '/onboarding', name: 'onboarding', component: OnboardingView },
    { path: '/wizard', name: 'wizard', component: WizardView },
    { path: '/login', name: 'login', component: LoginView },
    { path: '/paketler', name: 'paketler', component: WizardView }, // Temporary redirect
    { path: '/mobil', name: 'mobil', component: WizardView }, // Temporary redirect
    { path: '/ev-internet', name: 'ev-internet', component: WizardView }, // Temporary redirect
    { path: '/tv', name: 'tv', component: WizardView }, // Temporary redirect
    { path: '/kurumsal', name: 'kurumsal', component: WizardView }, // Temporary redirect
    { path: '/destek', name: 'destek', component: WizardView }, // Temporary redirect
    { path: '/api-test', name: 'api-test', component: ApiTestView },
    { path: '/:pathMatch(.*)*', redirect: '/onboarding' },
  ],
})
export default router
