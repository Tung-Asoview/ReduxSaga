import * as types from '../constants/actionTypes';

var initalState = [];

const artistReducer = (state = initalState, action) => {
    switch (action.type) {
        case types.ARTIST_ALL:
            state = [...action.artists];
            return [...state];
        // case ADD_MOVIE:
        //     return [
        //         ...state,
        //         action.newMovie
        //     ];
        default:
            return [...state]; //state does not change
    }
}

export default artistReducer;