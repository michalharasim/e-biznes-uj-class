package controllers

import play.api.libs.json._
import javax.inject._
import play.api._
import play.api.mvc._
import scala.collection.mutable
import models.Product
import models.ProductDTO._
import models.ProductDTO


@Singleton
class ProductController @Inject()(val controllerComponents: ControllerComponents) extends BaseController {
  private val productList = new mutable.ListBuffer[Product]()
  productList += Product(1, "Telefon", 1999.99)
  productList += Product(2, "Laptop", 4999.00)
  implicit val productListJson: OFormat[Product] = Json.format[Product]

  def getAll(): Action[AnyContent] = Action {
    if (productList.isEmpty) {
      NoContent
    } else {
      Ok(Json.toJson(productList))
    }
  }

  def findById(id: Long) = Action {
    val product = productList.find(product => product.id == id)
    product match {
      case Some(item) => Ok(Json.toJson(item))
      case None => NotFound(Json.obj("message" -> s"Product with ID $id not found."))

    }
  }

  def deleteById(id: Long): Action[AnyContent] = Action {
    val productToDelete = productList.find(_.id == id)

    productToDelete match {
      case Some(product) =>
        productList -= product
        Ok(Json.obj("message" -> s"Product with ID ${product.id} was deleted."))
      case None =>
        NotFound(Json.obj("message" -> s"Product with ID $id not found."))
    }
  }

  def createProduct(): Action[JsValue] = Action(parse.json) { request =>
    val content = request.body
    val jsonObject = Some(content)
    val newProduct: Option[ProductDTO] =
      jsonObject.flatMap(
        Json.fromJson[ProductDTO](_).asOpt
      )
    newProduct match {
      case Some(newItem) =>
        val nextId = productList.map(_.id).max + 1
        val toBeAdded = Product(nextId, newItem.name, newItem.price)
        productList += toBeAdded
        Created(Json.toJson(toBeAdded))
      case None =>
        BadRequest
    }
  }

  def updateProduct(id: Long): Action[JsValue] = Action(parse.json) { request =>
    request.body.validate[ProductDTO].fold(
      errors => BadRequest(Json.obj("error" -> "Invalid JSON format")),
      dto => {
        productList.find(_.id == id) match {
          case Some(existing) =>
            val updatedProduct = Product(id, dto.name, dto.price)
            productList -= existing
            productList += updatedProduct
            Ok(Json.obj("message" -> s"Product with ID $id updated."))
          case None =>
            NotFound
        }
      }
    )
  }
}
