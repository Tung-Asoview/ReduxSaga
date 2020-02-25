import * as types from '../constants/actionTypes';

var initalState = [];

const foodReducer = (state = initalState, action) => {
    switch (action.type) {
        case types.LIST_ALL:
            state = [...action.foods];
            return [...state];
        case types.FOODSBY:
            state = [...action.foodsBy]
            return [
                [...state]
            ];
        default:
            return [...state]; //state does not change
    }
}

export default foodReducer;