import {Site, SiteResponse} from "@/types/Site.ts";
import {alovaBaseUrlInstance} from "@/utils/http.ts";

export const useSitesStore = defineStore('sites', () => {
    const sites = ref<Site[]>([])
    const loadSites = async () => {
        const res: SiteResponse = await alovaBaseUrlInstance.Get('/api/sites')
        if (res.code === 200) {
            sites.value = res.data
        }
    }
    const deleteSite = async (id: number) => {
        const res: SiteResponse = await alovaBaseUrlInstance.Delete(`/api/site/${id}`)
        if (res.code === 200) {
            sites.value = sites.value.filter(site => site.ID !== id)
        }
    }
    return {
        sites,
        loadSites,
        deleteSite
    }
})