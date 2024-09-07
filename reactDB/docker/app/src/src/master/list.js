var React = require('react');
var DateFormat = require('dateformat');

const util = require('util')

class MasterTable extends React.Component {

  constructor(props) {

    super(props);

    console.log("props="+util.inspect(this.props));

    var clicked = false;
    var is_edit = false;

    if (this.props.is_edit) {
      clicked = true;
      is_edit = true;
    }

    this.state = {
      clicked: clicked,
      is_edit: is_edit,
    };
    this.handleClick = this.handleClick.bind(this);
  }

  handleClick() {

    if (this.state.clicked) {
      this.deleteMaster();
    } else {
      this.editMaster();
    }

    this.setState({
      clicked: !this.state.clicked,
      is_edit: false
    });

  }

  editMaster(props) {

    console.log("editMaster", this.props.name);
     return fetch("/_api/master/edit/"+this.props.name+"/"+this.props.user.displayName)
      .then((response) => response.json())
      .then((responseJson) =>
        this.setState({
          status: responseJson.status,
        })
      )
      .catch((error) => {
        console.error(error);
      });
  }

  deleteMaster(props) {
    console.log("deleteMaster", this.props.name);
     return fetch("/_api/master/delete/"+this.props.name)
      .then((response) => response.json())
      .then((responseJson) =>
        this.setState({
          status: responseJson.status,
        })
      )
      .catch((error) => {
        console.error(error);
      });
  }


  render() {

    var edit_time = "";
    var background = "";
    var editor = "";

    if (this.state.clicked) {
      background = "#ffb6c1";

      if (this.state.is_edit) {
        edit_time = this.props.edit_time;
        editor = this.props.editor;
      } else {
        var now = new Date();
        edit_time = DateFormat(now, "yyyy/mm/dd HH:MM:ss");
        editor = this.props.user.displayName;
      }
    } else {
      background = "#dbebf8";
    }

    return <li class="table_list" style={{ background }} onClick={this.handleClick}><div class="table_name">{this.props.name}</div> <div class="editor">{editor}</div> <div class="edit_time">{edit_time}</div></li>;
  }

}



class MasterList extends React.Component {


  constructor(props) {
    super(props);

    this.state = {
      tables: {},
    };
    this.loadMasterList = this.loadMasterList.bind(this);
  }


  loadMasterList() {
    return fetch("/_api/master/")
      .then((response) => response.json())
      .then((responseJson) =>
        this.setState({
          tables: responseJson.tables,
        })
      )
      .catch((error) => {
        console.error(error);
      });
  }

  componentWillMount() {
    this.loadMasterList();
  }


  render() {

    var master_list = [];
    for (var k in this.state.tables) {
      master_list.push(<MasterTable key={k} name={k} is_edit={this.state.tables[k].is_edit} editor={this.state.tables[k].editor} edit_time={this.state.tables[k].edit_time} user={this.props.user}/>);
    }

    return(
      <ul class="master_list">
        {master_list}
      </ul>
    );
  }
}


module.exports = MasterList;
