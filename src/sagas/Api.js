/*
Mr Nguyen Duc Hoang
https://www.youtube.com/c/nguyenduchoang
Email: sunlight4d@gmail.com
Send GET / POST api requests to server
*/
const apiGetAllFoods = 'https://facebook.github.io/react-native/movies.json';

async function getFoods() {
    try {
        let response = await fetch(apiGetAllFoods);
        let responseJson = await response.json();
        return responseJson.movies; //list of foods
    } catch (error) {
        console.error(`Error is : ${error}`);
    }
}

export default getFoods;