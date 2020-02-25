import React from "react";
import { 
  View,
  Text,
  StyleSheet,
  FlatList,
  Image,
  TouchableOpacity,
  TextInput,
  RefreshControl
} from "react-native";
import LinearGradient from "react-native-linear-gradient";
import AntDesign from "react-native-vector-icons/AntDesign";
import Ionicons from "react-native-vector-icons/Ionicons";
import { connect } from 'react-redux';
import { getFoods } from './../../sagas/Api';
import { listAll } from './../../actions/index';

class All extends React.Component{
  constructor(props){
    super(props);
    this.state={
      data: [],
      data_temp: [],
      refreshing: false,
      search: ''
    }
  }

  getAll = () => {
    this.setState({
      refreshing: true
    })
    getFoods().then(foods => {
      this.props.listAll(foods);
      this.setState({
        data: foods,
        data_temp: foods,
        refreshing: false
      })
    }).catch((error) => {
      this.setState({
        data: [],
        data_temp: [],
        refreshing: false
      })
    });
  }

  UNSAFE_componentWillMount() {
      this.getAll();
  }

  _rating(item){
      let rating = [];
      for(i=0;i<item;i++){
        rating.push(
          <Image 
            source={require("../../asset/star.png")}
            style={{width:15, height:15, marginRight:3}}
            resizeMode="cover"
          />
        )
      }
      return rating;
  }

  renderItem = ({item}) => {
    return(
        <LinearGradient 
        colors={['#009245', '#8cc631']}
        start={{x:0, y:1}} end={{x:1, y:0}}
        style={styles.item}
        >
          <View style={styles.image_container}>
              <Image 
                source={{uri: item.image}}
                style={styles.image}
              />
          </View>
          <View style={styles.content}>
              <Text style={styles.name}>{item.name}</Text>
              <View style={styles.rating}>
                {this._rating(item.avgStars)}
              </View>
              <View style={styles.price_container}>
                  <View style={styles.price}>
                      <Text style={styles.textPrice}>{item.price}</Text>
                  </View>
              </View>
          </View>
          <TouchableOpacity 
          onPress={()=>this.props.props.navigation.navigate("DetailScreen",{
            item: item
          })}
          style={styles.button}>
              <AntDesign 
                name="arrowright"
                color="green"
                size={15}
              />
          </TouchableOpacity>

        </LinearGradient>
    )
  }

  ItemSeparatorComponent = () => {
    return(
      <View 
        style={{
          height:10
        }}
      />
    )
  }

  _search(text){
      let data = this.state.data_temp.filter(value => {
        return value.title.toUpperCase().indexOf(text) !== -1;
      });
      this.setState({
        data:data,
        search:text
      });
  }

  render(){
    return(
      <View style={styles.container}>
          <View style={styles.section}>
              <TextInput 
                placeholder="Search.."
                style={{ flex:1, marginLeft:10}}
                value={this.state.search}
                keyboardAppearance={true}
                onChangeText={(text)=>this._search(text)}
              />
              <TouchableOpacity
              onPress={()=>this._search("")}
              style={{paddingHorizontal:10}}>
                <Ionicons 
                  name="ios-close"
                  color="gray"
                  size={20}
                />
              </TouchableOpacity>

          </View>
          <View style={styles.flatList}>
              <FlatList
                data={this.state.data}
                renderItem={this.renderItem}
                keyExtractor={(item, index)=>index.toString()}
                ItemSeparatorComponent={this.ItemSeparatorComponent}
                showsVerticalScrollIndicator={true}
                refreshControl={
                <RefreshControl
                  refreshing={this.state.refreshing}
                  onRefresh={this.getAll}
                />
              }
              />
          </View>
      </View>
    )
  }
}

var styles = StyleSheet.create({
  container: {
    flex:1,
    backgroundColor:'white',
    paddingBottom:5
  },
  flatList: {
    flex:1,
    marginTop:10
  },
  item: {
    flex:1,
    paddingVertical:10,
    paddingHorizontal:10,
    flexDirection:'row',
    borderRadius:10
  },
  image_container: {
    width: 90,
    height: 90
  },
  image: {
    width:'100%',
    height:'100%',
    borderWidth:5,
    borderColor:'white',
    borderRadius:10
  },
  content: {
    flex:1,
    justifyContent:'center',
    paddingHorizontal:10
  },
  name: {
    color:'white',
    fontWeight:'bold',
    fontSize:18
  },
  rating: {
    marginTop:5,
    flexDirection:'row'
  },
  button: {
    width:30,
    height:30,
    backgroundColor:'white',
    borderRadius:15,
    justifyContent:'center',
    alignItems:'center'
  },
  price_container: {
    flexDirection:'row',
    marginTop:10
  },
  price: {
    backgroundColor:'white',
    paddingVertical:5,
    paddingHorizontal:15,
    borderRadius:50
  },
  textPrice: {
    color:'green',
    fontWeight:'bold'
  },
  section: {
    flexDirection:'row',
    alignItems:'center',
    paddingVertical:5,
    paddingHorizontal:10,
    borderRadius:100,
    backgroundColor:'#f2f2f2',
    marginTop:10
  }
});

function mapStateToProps(state) {
  return { 
    foods: state.foods
  };
}

function mapDispatchToProps(dispatch, props) {
  return {
    listAll : (foods) => {
      dispatch(listAll(foods));
    }
  }
}

export default connect(mapStateToProps, mapDispatchToProps)(All);
