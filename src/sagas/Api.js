const apiFoods = 'http://localhost:8080/product/show/';
const apiArtists = 'http://localhost:8080/artist/show/';

export async function getFoods() {
    try {
        let response = await fetch(apiFoods);
        let responseJson = await response.json();
        return responseJson; //list of foods
    } catch (error) {
        console.error(`Error is : ${error}`);
    }
}

export async function getFoodByArtist(id) {
    try {
        let response = await fetch(apiFoods + '/artist/' + id);
        let responseJson = await response.json();
        return responseJson; //food
    } catch (error) {
        console.error(`Error is : ${error}`);
    }
}

export async function getArtists() {
    try {
        let response = await fetch(apiArtists);
        let responseJson = await response.json();
        return responseJson; //list of artists
    } catch (error) {
        console.error(`Error is : ${error}`);
    }
}