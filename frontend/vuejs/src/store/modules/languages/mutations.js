import i18next from 'i18next'
import { languageTypes } from '../../mutation-types'

export default {
  [languageTypes.CHANGE_LOCALE] (state, locale) {
    state.locale = locale
    i18next.changeLanguage(locale)
  }
}
