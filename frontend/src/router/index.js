import { createRouter, createWebHistory } from 'vue-router'
import WizardView from '../views/WizardView.vue'
import LoginView from '../views/LoginView.vue'
import OnboardingView from '../views/OnboardingView.vue'
import CoverageView from '../views/CoverageView.vue'
import RecommendationsView from '../views/RecommendationsView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    { path: '/', redirect: '/onboarding' },
    { path: '/onboarding', name: 'onboarding', component: OnboardingView },
    { path: '/wizard', name: 'wizard', component: WizardView },
    {
      path: '/wizard/step1',
      name: 'wizard-step1',
      component: WizardView,
      props: { initialStep: 'form' },
    },
    {
      path: '/wizard/step2',
      name: 'wizard-step2',
      component: WizardView,
      props: { initialStep: 'recommendations' },
    },
    {
      path: '/wizard/step3',
      name: 'wizard-step3',
      component: WizardView,
      props: { initialStep: 'coverage' },
    },
    { path: '/login', name: 'login', component: LoginView },
    { path: '/coverage', name: 'coverage', component: CoverageView },
    { path: '/recommendations', name: 'recommendations', component: RecommendationsView },
    { path: '/paketler', name: 'paketler', component: WizardView }, // Temporary redirect
    { path: '/mobil', name: 'mobil', component: WizardView }, // Temporary redirect
    { path: '/ev-internet', name: 'ev-internet', component: WizardView }, // Temporary redirect
    { path: '/tv', name: 'tv', component: WizardView }, // Temporary redirect
    // { path: '/kurumsal', name: 'kurumsal', component: WizardView }, // Removed - now uses #
    // { path: '/destek', name: 'destek', component: WizardView }, // Removed - now uses #
    { path: '/:pathMatch(.*)*', redirect: '/onboarding' },
  ],
})
export default router