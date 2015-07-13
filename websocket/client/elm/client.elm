import Html exposing (..)
import Html.Attributes exposing (..)
import Html.Events exposing (on, targetChecked)
import Signal exposing (Address)
import List
import StartApp

main =
  StartApp.start { model = initialModel, view = view, update = update }


-- MODEL

type alias Model =
  { history: List String
  }


initialModel =
  {history = ["hello"]}


-- UPDATE

type Action
  = Add String


update action model =
  case action of
    Add string ->
        { model | history <- string :: model.history }

-- VIEW

view address model =
    div [] <| List.map text model.history ++ [input
                                              [ placeholder "Type text"
                                              , on "input" targetValue (\str -> Signal.message address (Add str))
                                              ] []]