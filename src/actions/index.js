import * as types from '../constants/actionTypes';

export const listAll = (foods) => {
    return {
        type: types.LIST_ALL,
        foods
    };
}

export const foodsByArtistId = (foodsBy) => {
    return {
        type: types.FOODSBY,
        foodsBy
    };
}

export function artistAll(artists) {
    return { 
        type: types.ARTIST_ALL,
        artists
    };
}

// export function fetchError() {
//     return { type: 'FETCH_ERROR' };
// }

// export function fetchDataThunk() {
//     return dispatch => {
//         dispatch(fetchSuccess(foods));
//         getFoods()
//         .then(res => fetchSuccess(res))
//     };
// }