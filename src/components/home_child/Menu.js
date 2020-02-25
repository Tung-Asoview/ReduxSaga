import React from "react";
import { 
  View,
  Text,
  StyleSheet,
  FlatList,
  Image,
  TouchableOpacity
} from "react-native";
import { connect } from 'react-redux';
import { getFoodByArtist, getArtists } from './../../sagas/Api';
import { foodsByArtistId, artistAll } from './../../actions/index';

var test = [
  {
     "changedBy":1,
     "changedOn":"2018-12-31T17:00:00.000+0000",
     "country":"Việt Nam",
     "createdBy":1,
     "createdOn":"2018-12-31T17:00:00.000+0000",
     "description":"Sản xuất tại việt nam",
     "firstName":"Kho hàng",
     "id":1,
     "lastName":"Quần áo",
     "lifeSpan":"LifeSpan",
     "totalProducts":100
  },
  {
     "changedBy":1,
     "changedOn":"2018-12-31T17:00:00.000+0000",
     "country":"Hàn Quốc",
     "createdBy":1,
     "createdOn":"2018-12-31T17:00:00.000+0000",
     "description":"Sản xuất tại Hàn Quốc",
     "firstName":"Kho hàng",
     "id":2,
     "lastName":"Quần áo",
     "lifeSpan":"LifeSpan",
     "totalProducts":100
  },
  {
     "changedBy":1,
     "changedOn":"2018-12-31T17:00:00.000+0000",
     "country":"Nhật bản",
     "createdBy":1,
     "createdOn":"2018-12-31T17:00:00.000+0000",
     "description":"Sản xuất tại Nhật bản",
     "firstName":"Kho hàng",
     "id":3,
     "lastName":"Quần áo",
     "lifeSpan":"LifeSpan",
     "totalProducts":100
  },
  {
     "changedBy":1,
     "changedOn":"2018-12-31T17:00:00.000+0000",
     "country":"Singapore",
     "createdBy":1,
     "createdOn":"2018-12-31T17:00:00.000+0000",
     "description":"Sản xuất tại Singapore ,việt nam",
     "firstName":"Kho hàng",
     "id":4,
     "lastName":"Quần áo",
     "lifeSpan":"LifeSpan",
     "totalProducts":100
  },
  {
     "changedBy":1,
     "changedOn":"2018-12-31T17:00:00.000+0000",
     "country":"Philipin",
     "createdBy":1,
     "createdOn":"2018-12-31T17:00:00.000+0000",
     "description":"Sản xuất tại Philipin",
     "firstName":"Kho hàng",
     "id":5,
     "lastName":"Quần áo",
     "lifeSpan":"LifeSpan",
     "totalProducts":100
  }
]

class Menu extends React.Component{

  constructor(props){
    super(props);
    this.state={
      data: [],
      foods: props.foodsBy
    }
  }
  getAll(){
    getArtists().then(artists => {
      // console.log(artists)
      this.props.artistAll(artists);
      this.setState({
        data: artists
      })
    });
  }

  getFoods(id){
    getFoodByArtist(id).then(foods => {
      this.props.foodsByArtistId(foods);
    })
  }

  componentDidMount() {
    this.getAll();
  }

  renderItem_type = ({item}) => {
    return(
        <TouchableOpacity 
        onPress={()=>this.props.props.navigation.navigate("DetailScreen",{
          image: item.image,
          price: item.price,
          name: item.name
        })}
        style={styles.item_type}>
            <Image 
              source={item.image}
              style={styles.image}
            />
            <Text style={styles.name}>{item.name}</Text>
        </TouchableOpacity>

    )
  }

  renderItem = ({item}) => {
    console.log(this.state.foods)
    // this.getFoods(item.id);
    return(
      <View style={{flex:1}}>
          <Text style={[styles.type,{
            color: '#f7931e'
          }]}>{item.description}</Text>
          <View style={[styles.item,{
            backgroundColor:'#f7931e'
          }]}>
              <FlatList 
                data={this.props.foodsBy}
                renderItem={this.renderItem_type}
                keyExtractor={(item,index)=>index.toString()}
                horizontal={true}
                showsHorizontalScrollIndicator={false}
                ItemSeparatorComponent={this.ItemSeparatorComponent_type}
              />
          </View>
      </View>
    )
  }

  ItemSeparatorComponent_type = () => {
    return(
      <View 
        style={{width:10}}
      />
    )
  }
  
  ItemSeparatorComponent = () => {
    return(
      <View 
        style={{height:20}}
      />
    )
  }
  
  render(){
    console.log(this.state.data)
    return(
      <View style={styles.container}>
          <FlatList 
            data={this.state.data}
            renderItem={this.renderItem}
            keyExtractor={(item,index)=>index.toString()}
            showsVerticalScrollIndicator={false}
            ItemSeparatorComponent={this.ItemSeparatorComponent}
          />
      </View>
    )
  }
}

var styles = StyleSheet.create({
  container: {
    flex:1,
    marginTop:10,
    marginBottom:10,
    backgroundColor:'white'
  },
  type: {
    fontSize:18,
    fontWeight:'bold'
  },
  item: {
    marginTop:10,
    flexDirection:'row',
    paddingHorizontal:15,
    paddingVertical:10,
    borderRadius:10
  },
  item_type: {
    flex:1,
    alignItems:'center'
  },
  image:{
    width:80,
    height:80,
    borderRadius:40,
    borderWidth:2,
    borderColor:'white'
  },
  name: {
    marginTop:10,
    color:'white',
    fontSize:15
  }
});

function mapStateToProps(state) {
  return { 
    artists : state.artists,
    foodsBy: state.foodsBy
  };
}

function mapDispatchToProps(dispatch, props) {
  return {
    artistAll : (artists) => {
      dispatch(artistAll(artists));
    },
    foodsByArtistId : (foodsBy) => {
      dispatch(foodsByArtistId(foodsBy));
    }
  }
}

export default connect(mapStateToProps, mapDispatchToProps)(Menu);