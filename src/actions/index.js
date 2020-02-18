import * as types from '../constants/actionTypes';
import getFoods from '../sagas/Api';

export const listAll = () => {
    return {
        type: types.LIST_ALL
    }
}

export function fetchSuccess(foods) {
    return { 
        type: 'FETCH_SUCCESS',
        cityName,
        temp 
    };
}

export function fetchError() {
    return { type: 'FETCH_ERROR' };
}

export function fetchDataThunk() {
    return dispatch => {
        dispatch(startFetchData());
        getFoods()
        .then(temp => dispatch(fetchSuccess(foods)))
        .catch(() => dispatch(fetchError()));
    };
}