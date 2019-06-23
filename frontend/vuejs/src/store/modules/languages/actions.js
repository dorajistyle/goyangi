import {languageTypes} from '../../mutation-types'
export default {
  changeLocale ({ commit }, locale) {
    console.log('locale to : '+ locale)
    commit(languageTypes.CHANGE_LOCALE, locale)
  }
}
