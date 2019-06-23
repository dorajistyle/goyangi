import Articles from '@/views/articles'
import ArticleList from '@/components/articles/List'
import ArticleItem from '@/components/articles/Item'
import ArticleForm from '@/components/articles/Form'

export default [
  {
    path: '/articles',
    component: Articles,
    children: [
      {
        path: '',
        component: ArticleList,
        props: (route) => ({ currentPage: route.query.currentPage })
      },
      {
        path: 'new',
        component: ArticleForm,
        meta: { requiresAuth: true },
        props: {action: 'new'}
      },
      {
        path: ':articleId',
        component: ArticleItem,
        props: true
      },
      {
        path: ':articleId/edit',
        component: ArticleForm,
        props: {action: 'edit'}
      }
    ]
  }
]
