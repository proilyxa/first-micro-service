import type {VueElement} from 'vue'
import type {AppLayoutsEnum} from '@/layouts/layouts.types'
import 'vue-router'

export {}

declare module 'vue-router' {
    interface RouteMeta {
        layout?: AppLayoutsEnum
        layoutComponent?: VueElement
    }
}