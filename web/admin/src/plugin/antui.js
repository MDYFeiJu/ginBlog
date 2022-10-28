import {
  Button,
  Form,
  Input,
  Layout,
  Menu,
  Card,
  Table,
  Row,
  Col,
} from 'ant-design-vue'
const components = [Button, Form, Input, Layout, Menu, Card, Table, Row, Col]

export default function setupAtnd(app) {
  components.forEach((component) => {
    app.use(component)
  })
}
