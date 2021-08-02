import React, { Component } from "react";
import { api } from  "../js/api.js"
import Badge from 'react-bootstrap/Badge'
import "bootstrap/dist/css/bootstrap.css";
import { Button, Modal, ModalHeader, ModalBody, ModalFooter } from "reactstrap";
import ReactMarkdown from "react-markdown";
import Tab from 'react-bootstrap/Tab'
import Nav from 'react-bootstrap/Nav'
import TabContainer from 'react-bootstrap/TabContainer'
import TabContent from 'react-bootstrap/TabContent'
import TabPane from 'react-bootstrap/TabPane'

class Challenges extends Component {

	// openModal = () => this.setState({ isOpen: true });
	// closeModal = () => this.setState({ isOpen: false });


	render() {
		
		return (
			<React.Fragment>
				<div className="container">
					<div className="row mt-5">
						<h2>Scoreboard</h2>
					</div>
					<br />
					<h3>CTF has not started</h3>
				</div>
			</React.Fragment>
		);
	}
}

export default Challenges;
