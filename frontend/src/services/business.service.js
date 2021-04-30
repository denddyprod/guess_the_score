import axios from "axios";

const API_URL = "http://localhost:5000/";

class BusinessService {
    allMatches() {
        return axios
            .get(API_URL + "matches")
            .then(response => {
                console.log(response.data)
                return response.data;
            });
    }

    // getCurrentUser() {
    //     return JSON.parse(localStorage.getItem('user'));;
    // }
}

export default new BusinessService();